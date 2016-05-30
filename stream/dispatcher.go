package stream

import "time"
import "math"

// Dispatcher sends stream data to an accumulator based on ordinal value
type Dispatcher struct {
    key                     string
    intervalSize            int64
    accumulatorMap          map[int64]*Accumulator
    targetSampleCount       int
}

// NewDispatcher creates a new dispatcher
func NewDispatcher(key string, intervalSize int64) (*Dispatcher) {
    var accumulatorMap = make(map[int64]*Accumulator)
    
    dispatcher := Dispatcher {
        key: key,
        intervalSize: intervalSize,
        accumulatorMap: accumulatorMap,
        targetSampleCount: 10000}
    
    return &dispatcher
}

// IncludeForOrdinal includes a new value for an ordinal value
func (dispatcher *Dispatcher) IncludeForOrdinal(ordinal int64, value float64) {
    var interval = int64(math.Floor(float64(ordinal) / float64(dispatcher.intervalSize)))
    
    // lookup the appropriate accumulator based on interval
    accumulator := dispatcher.accumulatorMap[interval]
    
    if accumulator == nil {
        accumulator = NewAccumulator(dispatcher.targetSampleCount)
        
        dispatcher.accumulatorMap[interval] = accumulator
    }
    
    accumulator.Include(value)
}

// IncludeForTime includes a new value, dispatching it to the appropirate accumulator based on time interval
func (dispatcher *Dispatcher) IncludeForTime(t time.Time, value float64) {
    dispatcher.IncludeForOrdinal(t.UnixNano(), value)
}

// FinaliseIntervalsPriorToOrdinal causes all accumulators for intervals prior to
// that related to the specified ordinal and removes them from the dispatcher
func (dispatcher *Dispatcher) FinaliseIntervalsPriorToOrdinal(ordinal int64) {
    var interval = int64(math.Floor(float64(ordinal) / float64(dispatcher.intervalSize)))
    
    for k,v := range dispatcher.accumulatorMap {
        if k < interval {
            // finalise and remove the accumulator
            var _ = v.Finalise()
            delete(dispatcher.accumulatorMap, k)
        }
    }
}

// FinaliseIntervalsPriorToTime causes all accumulators for intervals prior to
// that related to the specified time, and removes them from the dispatcher
func (dispatcher *Dispatcher) FinaliseIntervalsPriorToTime(t time.Time) {
    dispatcher.FinaliseIntervalsPriorToOrdinal(t.UnixNano())
}
