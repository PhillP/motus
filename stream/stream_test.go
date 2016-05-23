package stream

import "testing"
import "github.com/stretchr/testify/assert"
import "fmt"

// Test that the Accumulator produces correct results when no values are received within the interval
func TestAccumulatorCycle_NoValue(t *testing.T) {
    var targetSampleCount = 100 // the sample count is larger than the number of values provided within this test
    
    var accumulator = NewAccumulator(targetSampleCount)
    
    // this test covers a scenario in which no values are received within the interval
    // .. no values are provided to the accumulator ..
        
    // finalise the accumulator and gather statistics
    var statistics = accumulator.Finalise()
    
    fmt.Printf("Count is %d", accumulator.count)
    assert.Equal(t, int64(0), statistics.count, "Count should be 0 when no values where provided to the accumulator")
    assert.Equal(t, statistics.count, int64(statistics.sampleCount), "Sample count must count when all values are included in the sample set")
}

// Test that the Accumulator produces correct results when only a single value is received within the interval
func TestAccumulatorCycle_SingleValue(t *testing.T) {
    var targetSampleCount = 100 // the sample count is larger than the number of values provided within this test
    
    var accumulator = NewAccumulator(targetSampleCount)
    
    var value float64 = 50
    
    // this test covers a scenario in which only 1 value is received within the interval
    accumulator.Include(float64(value))
    
    // finalise the accumulator and gather statistics
    var statistics = accumulator.Finalise()
    
    fmt.Printf("Count is %d", accumulator.count)
    
    assert.Equal(t, int64(1), statistics.count, "Count should be 1 when a single value was provided to the accumulator")
    assert.Equal(t, value, statistics.mean, "Mean should equal the single value when only a single value was provided to the accumulator")
    assert.Equal(t, value, statistics.minimum, "Minimum should equal the single value when only a single value was provided to the accumulator")
    assert.Equal(t, value, statistics.maximum, "Maximum should equal the single value when only a single value was provided to the accumulator")
    assert.Equal(t, float64(0), statistics.sampleStandardDeviation, "Standard deviation should be 0 when only a single value was provided")
    assert.Equal(t, float64(0), statistics.coefficientOfVariation, "Coefficient of variation must be 0 when standard deviation is 0")
    
    assert.Equal(t, statistics.mean, statistics.sampleMean, "Sample mean must equal mean when all values are included in the sample set")
    assert.Equal(t, statistics.sum, statistics.sampleSum, "Sample sum must equal sum when all values are included in the sample set")
    assert.Equal(t, statistics.count, int64(statistics.sampleCount), "Sample count must count when all values are included in the sample set")
}

// Test that the Accumulator produces correct results when multiple values are received (and all included in the sample count) within the interval
func TestAccumulatorCycle_SomeValues(t *testing.T) {
    assert.Fail(t,"Test not implemented")
}

// Test that the Accumulator produces correct results when multiple values (where some are excluded from the sample count) are received within the interval
func TestAccumulatorCycle_ManyValues(t *testing.T) {
    assert.Fail(t,"Test not implemented")
}
