package stream

import (
 "time"
 "math"
)

// IntervalRouter sends stream data to an accumulator based matching the ordinal value to an interval 
type IntervalRouter struct {
    key                     string
    intervalSize            int64
    intervalType            IntervalType
    maxIntervalLag          uint32
    maxInterval             int64
    channelMap              map[int64]chan OrdinalValue
    doneMap                 map[int64]chan bool
    targetSampleCount       uint32
}

// NewIntervalRouter creates a new router used to assign values to intervals
func NewIntervalRouter(key string, 
                intervalSize int64, 
                intervalType IntervalType,
                maxIntervalLag uint32,
                targetSampleCount uint32) (*IntervalRouter) {
    var channelMap = make(map[int64]chan OrdinalValue)
    var doneMap = make(map[int64]chan bool)
    
    intervalRouter := IntervalRouter {
        key: key,
        intervalSize: intervalSize,
        intervalType: intervalType,
        maxIntervalLag: maxIntervalLag,
        maxInterval: math.MinInt64, // set to the minimum value rather than 0 as negative intervals are supported
        channelMap: channelMap,
        doneMap: doneMap,
        targetSampleCount: targetSampleCount}
    
    return &intervalRouter
}

// AccumulateFromChannel directs a value to the appropriate accumulator for an ordinal value
func (intervalRouter *IntervalRouter) AccumulateFromChannel(input chan OrdinalValue, output chan IntervalStatistics, done chan bool) {
    for v := range input {
        intervalRouter.Accumulate(v, output)
    }
    
    intervalRouter.FinaliseAll()
    
    done <- true
}

// Accumulate directs a value to the appropriate accumulator for an ordinal value
func (intervalRouter *IntervalRouter) Accumulate(ordinalValue OrdinalValue, output chan IntervalStatistics) {
    var interval = int64(math.Floor(float64(ordinalValue.Ordinal) / float64(intervalRouter.intervalSize)))
    
    // only process ordinals within the lag range
    if intervalRouter.maxInterval == math.MinInt64 || interval >= (intervalRouter.maxInterval - int64(intervalRouter.maxIntervalLag)) {
        
        // lookup the appropriate channel based on interval
        channel := intervalRouter.channelMap[interval]
        
        var intervalStart = interval * int64(intervalRouter.intervalSize)
        
        if channel == nil {
            accumulator := NewAccumulator(intervalStart, intervalStart + intervalRouter.intervalSize - 1, intervalRouter.intervalType, intervalRouter.targetSampleCount)
            
            channel = make(chan OrdinalValue)
            doneChannel := make(chan bool)
        
            intervalRouter.channelMap[interval] = channel
            intervalRouter.doneMap[interval] = doneChannel
            
            // start accumulating
            go accumulator.Accumulate(channel, output, doneChannel)
        }
        
        channel <- ordinalValue
        
        if intervalRouter.maxInterval == math.MinInt64 {
            intervalRouter.maxInterval = interval
        } else if interval > intervalRouter.maxInterval {
            // max interval increased
            intervalRouter.maxInterval = interval
            
            var maxIntervalToKeep = intervalRouter.maxInterval - int64(intervalRouter.maxIntervalLag)
            var maxOrdinalToKeep = maxIntervalToKeep * intervalRouter.intervalSize
            
            intervalRouter.FinalisePriorTo(maxOrdinalToKeep)
        }
    }
}

// FinalisePriorTo causes all accumulators for intervals prior to
// that related to the specified ordinal and removes them from the intervalRouter
func (intervalRouter *IntervalRouter) FinalisePriorTo(ordinal int64) {
    var interval = int64(math.Floor(float64(ordinal) / float64(intervalRouter.intervalSize)))
    
    for k,c := range intervalRouter.channelMap {
        if k < interval {
            // close the channel
            close(c)
            
            // remove the accumulator
            delete(intervalRouter.channelMap, k)
        }
    }
    
    for k,c := range intervalRouter.doneMap {
        if k < interval {
            // wait for completion
            <- c
            
            // close the channel
            close(c)
            
            // remove the channel from the map
            delete(intervalRouter.doneMap, k)
        }
    }
}

// FinalisePriorToTime causes all accumulators for intervals prior to
// that related to the specified time, and removes them from the intervalRouter
func (intervalRouter *IntervalRouter) FinalisePriorToTime(t time.Time) {
    intervalRouter.FinalisePriorTo(t.UnixNano())
}

// FinaliseAll causes all accumulators to be finalised
func (intervalRouter *IntervalRouter) FinaliseAll() {
    intervalRouter.FinalisePriorTo(math.MaxInt64)
}
