package stream

import "time"
import "math"

// Dispatcher sends stream data to an accumulator based on ordinal value
type Dispatcher struct {
    key                     string
    intervalSize            int64
    intervalType            IntervalType
    channelMap              map[int64]chan OrdinalValue
    targetSampleCount       int
}

// NewDispatcher creates a new dispatcher
func NewDispatcher(key string, intervalSize int64, intervalType IntervalType) (*Dispatcher) {
    var channelMap = make(map[int64]chan OrdinalValue)
    
    dispatcher := Dispatcher {
        key: key,
        intervalSize: intervalSize,
        intervalType: intervalType,
        channelMap: channelMap,
        targetSampleCount: 10000}
    
    return &dispatcher
}

// AccumulateFromChannel directs a value to the appropriate accumulator for an ordinal value
func (dispatcher *Dispatcher) AccumulateFromChannel(input chan OrdinalValue, output chan IntervalStatistics) {
    for v := range input {
        dispatcher.Accumulate(v, output)
    }
    
    dispatcher.FinaliseAll()
}

// Accumulate directs a value to the appropriate accumulator for an ordinal value
func (dispatcher *Dispatcher) Accumulate(ordinalValue OrdinalValue, output chan IntervalStatistics) {
    var interval = int64(math.Floor(float64(ordinalValue.ordinal) / float64(dispatcher.intervalSize)))
    
    // lookup the appropriate channel based on interval
    channel := dispatcher.channelMap[interval]
    
    if channel == nil {
        accumulator := NewAccumulator(interval, interval + dispatcher.intervalSize, dispatcher.intervalType, dispatcher.targetSampleCount)
        
        channel = make(chan OrdinalValue)
        dispatcher.channelMap[interval] = channel
        
        // start accumulating
        go accumulator.Accumulate(channel, output)
    }
    
    channel <- ordinalValue
}

// FinalisePriorTo causes all accumulators for intervals prior to
// that related to the specified ordinal and removes them from the dispatcher
func (dispatcher *Dispatcher) FinalisePriorTo(ordinal int64) {
    var interval = int64(math.Floor(float64(ordinal) / float64(dispatcher.intervalSize)))
    
    for k,v := range dispatcher.channelMap {
        if k < interval {
            // close the channel
            close(v)
            
            // remove the accumulator
            delete(dispatcher.channelMap, k)
        }
    }
}

// FinalisePriorToTime causes all accumulators for intervals prior to
// that related to the specified time, and removes them from the dispatcher
func (dispatcher *Dispatcher) FinalisePriorToTime(t time.Time) {
    dispatcher.FinalisePriorTo(t.UnixNano())
}

// FinaliseAll causes all accumulators to be finalised
func (dispatcher *Dispatcher) FinaliseAll() {
    dispatcher.FinalisePriorTo(math.MaxInt64)
}
