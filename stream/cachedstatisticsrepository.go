package stream

import "sync"

// CachedStatisticsRepository provides access to generated interval statistics 
type CachedStatisticsRepository struct {
    mu                  *sync.RWMutex
    cacheByKey          map[string]*IntervalStatisticsCache
}

// NewCachedStatisticsRepository creates an empty statistics repository
func NewCachedStatisticsRepository() CachedStatisticsRepository {
    var cacheMap = make(map[string]*IntervalStatisticsCache, 0)
    
    return CachedStatisticsRepository {
        mu: &sync.RWMutex {},
        cacheByKey: cacheMap}
}

// RegisterCache includes a new key and cache within the repository
func (cachedStatisticsRepository *CachedStatisticsRepository) RegisterCache(key string, cache *IntervalStatisticsCache) error {
    defer cachedStatisticsRepository.mu.Unlock()
    cachedStatisticsRepository.mu.Lock()
    
    if _,ok := cachedStatisticsRepository.cacheByKey[key]; ok {
        // the key already had an entry... remove it
        delete(cachedStatisticsRepository.cacheByKey, key)
    }
    
    cachedStatisticsRepository.cacheByKey[key] = cache
    
    return nil
}

// UnregisterCache removes a cache and key from the repository
func (cachedStatisticsRepository *CachedStatisticsRepository) UnregisterCache(key string) error {
    defer cachedStatisticsRepository.mu.Unlock()
    cachedStatisticsRepository.mu.Lock()
    
    if _,ok := cachedStatisticsRepository.cacheByKey[key]; ok {
        // the key already had an entry... remove it
        delete(cachedStatisticsRepository.cacheByKey, key)
    }
    
    return nil
}

// EmitIntervalStatisticsForKeys outputs matching interval statistics to a provided channel
func (cachedStatisticsRepository *CachedStatisticsRepository) EmitIntervalStatisticsForKeys(keys []string, fromOrdinal int64, untilOrdinal int64, output chan IntervalStatistics) error {
    defer cachedStatisticsRepository.mu.RUnlock()
    cachedStatisticsRepository.mu.RLock()
    
    for _,key := range keys {
        if cache, ok := cachedStatisticsRepository.cacheByKey[key]; ok {
            cache.EmitIntervalStatistics(fromOrdinal, untilOrdinal, output)
        }
    }
    
    return nil
}
