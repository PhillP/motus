package stream

import "sync"

// IntervalStatisticsCache stores a limited set of interval statistics for a stream 
type IntervalStatisticsCache struct {
    streamKey   string
    size        uint32
    
    cacheMutex  sync.Mutex
    cache       []IntervalStatistics
}

// NewIntervalStatisticsCache creates a new cache for interval statistics
func NewIntervalStatisticsCache(streamKey string, size uint32) *IntervalStatisticsCache {
    var cache = make([]IntervalStatistics, 0)
    
    return &IntervalStatisticsCache {
        streamKey: streamKey,
        size: size,
        cache: cache};
}

// ProcessAndForward stores values from an input channel and then forwards values to an output channel. Only the most recent set of results are stored
func (intervalStatisticsCache *IntervalStatisticsCache) ProcessAndForward(input chan IntervalStatistics, output chan IntervalStatistics) {
    for v := range input {
       intervalStatisticsCache.addToCache(v)
        
        // forward
        output <- v
    }
}

func (intervalStatisticsCache *IntervalStatisticsCache) addToCache(statistics IntervalStatistics){
    defer intervalStatisticsCache.cacheMutex.Unlock()
    intervalStatisticsCache.cacheMutex.Lock()
    
    var newCache = intervalStatisticsCache.cache
    // remove from the cache if required
    if len(newCache) >= int(intervalStatisticsCache.size) {
        // keep all but the first element
        newCache = newCache[1:]
    }
    newCache = append(newCache, statistics)
    intervalStatisticsCache.cache = newCache    
}

// GetLast returns the most recent statistics from the cache up to the specified maxCount
func (intervalStatisticsCache *IntervalStatisticsCache) GetLast(maxCount int) []IntervalStatistics {
    var cache = intervalStatisticsCache.cache
    
    if (len(cache) > maxCount) {
        // limit the set of data to be returned
        cache = cache[len(cache) - maxCount:]
    }
    
    return cache
}

// GetFromOrdinal returns the cached statistics that have an ordinal value equal or greater than the value provided
func (intervalStatisticsCache *IntervalStatisticsCache) GetFromOrdinal(fromOrdinal int64) []IntervalStatistics {
    var cache = intervalStatisticsCache.cache
    var selected = make([]IntervalStatistics,0)
    
    for _,v := range cache {
       if (v.IntervalStart >= fromOrdinal) {
           selected = append(selected, v)
       }
    }
    
    return selected
}

// GetOrdinalRange returns the cached statistics that exist between a range of ordinal values
func (intervalStatisticsCache *IntervalStatisticsCache) GetOrdinalRange(fromOrdinal int64, untilOrdinal int64) []IntervalStatistics {
    var cache = intervalStatisticsCache.cache
    var selected = make([]IntervalStatistics,100)
    
    for _,v := range cache {
       if (v.IntervalStart >= fromOrdinal && v.IntervalStart < untilOrdinal) {
           selected = append(selected, v)
       }
    }
    
    return selected
}

// EmitIntervalStatistics outputs matching interval statistics to a provided channel
func (intervalStatisticsCache *IntervalStatisticsCache) EmitIntervalStatistics(fromOrdinal int64, untilOrdinal int64, output chan IntervalStatistics) error {
    var cache = intervalStatisticsCache.cache
    
    for _,v := range cache {
       if (v.IntervalStart >= fromOrdinal && v.IntervalStart < untilOrdinal) {
           output <- v
       }
    }
    
    return nil
}