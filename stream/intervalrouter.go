package stream

import "time"
import "math"

// IntervalRouter sends stream data to an accumulator based matching the ordinal value to an interval 
type IntervalRouter struct {
    key                     string
    intervalSize            int64
    intervalType            IntervalType
    channelMap              map[int64]chan OrdinalValue
    targetSampleCount       int
}

// NewIntervalRouter creates a new router used to assign values to intervals
func NewIntervalRouter(key string, intervalSize int64, intervalType IntervalType) (*IntervalRouter) {
    var channelMap = make(map[int64]chan OrdinalValue)
    
    intervalRouter := IntervalRouter {
        key: key,
        intervalSize: intervalSize,
        intervalType: intervalType,
        channelMap: channelMap,
        targetSampleCount: 10000}
    
    return &intervalRouter
}

// AccumulateFromChannel directs a value to the appropriate accumulator for an ordinal value
func (intervalRouter *IntervalRouter) AccumulateFromChannel(input chan OrdinalValue, output chan IntervalStatistics) {
    for v := range input {
        intervalRouter.Accumulate(v, output)
    }
    
    intervalRouter.FinaliseAll()
}

// Accumulate directs a value to the appropriate accumulator for an ordinal value
func (intervalRouter *IntervalRouter) Accumulate(ordinalValue OrdinalValue, output chan IntervalStatistics) {
    var interval = int64(math.Floor(float64(ordinalValue.ordinal) / float64(intervalRouter.intervalSize)))
    
    // lookup the appropriate channel based on interval
    channel := intervalRouter.channelMap[interval]
    
    if channel == nil {
        accumulator := NewAccumulator(interval, interval + intervalRouter.intervalSize, intervalRouter.intervalType, intervalRouter.targetSampleCount)
        
        channel = make(chan OrdinalValue)
        intervalRouter.channelMap[interval] = channel
        
        // start accumulating
        go accumulator.Accumulate(channel, output)
    }
    
    channel <- ordinalValue
}

// FinalisePriorTo causes all accumulators for intervals prior to
// that related to the specified ordinal and removes them from the intervalRouter
func (intervalRouter *IntervalRouter) FinalisePriorTo(ordinal int64) {
    var interval = int64(math.Floor(float64(ordinal) / float64(intervalRouter.intervalSize)))
    
    for k,v := range intervalRouter.channelMap {
        if k < interval {
            // close the channel
            close(v)
            
            // remove the accumulator
            delete(intervalRouter.channelMap, k)
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