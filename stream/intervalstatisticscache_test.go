package stream

import "testing"
import "github.com/stretchr/testify/assert"

func TestProcessAndForward(t *testing.T) {
    var input = make(chan IntervalStatistics, 0)
    var output = make(chan IntervalStatistics, 100)
    var cacheSize = 10
    
    var cache = NewIntervalStatisticsCache("MyStream", uint32(cacheSize))
    
    go ProcessAndForward(cache, input, output)
    
    var itemsToAdd = cacheSize * 2
    
    // add items to input
    for i := 0; i < itemsToAdd; i++ {
        input <- IntervalStatistics { IntervalStart : int64(i), IntervalEnd : int64(i + 1) }
    }
    close(input)
    
    var readCount = 0
    
    // read items from output
    for s := range output {
        assert.Equal(t, int64(readCount), s.IntervalStart)        
        readCount++
        
        if readCount == cacheSize * 2 {
            close(output)
        }
    }
    
    var expectedOrdinal = cacheSize
    // get items from cache.. as many as possible
    readCount = 0
    for _,s := range GetFromOrdinal(cache, 0) {
        assert.Equal(t, int64(expectedOrdinal), s.IntervalStart)
        expectedOrdinal++
        readCount++
    }
    
    // the expected number of items should have been returned from the cache
    assert.Equal(t, int64(cacheSize), int64(readCount))
}