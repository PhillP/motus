package stream

import "testing"
import "github.com/stretchr/testify/assert"
import "math"

// Test that the Accumulator produces correct results when no values are received within the interval
func TestAccumulatorCycle_NoValue(t *testing.T) {
    var targetSampleCount = 100 // the sample count is larger than the number of values provided within this test
    
    var accumulator = NewAccumulator(targetSampleCount)
    
    // this test covers a scenario in which no values are received within the interval
    // .. no values are provided to the accumulator ..
        
    // finalise the accumulator and gather statistics
    var statistics = accumulator.Finalise()
    
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
    
    assert.Equal(t, int64(1), statistics.count, "Count should be 1 when a single value was provided to the accumulator")
    assert.Equal(t, value, statistics.mean, "Mean should equal the single value when only a single value was provided to the accumulator")
    assert.Equal(t, value, statistics.minimum, "Minimum should equal the single value when only a single value was provided to the accumulator")
    assert.Equal(t, value, statistics.maximum, "Maximum should equal the single value when only a single value was provided to the accumulator")
    assert.Equal(t, value, statistics.sum, "Sum should equal the single value when only a single value was provided to the accumulator")
    assert.Equal(t, float64(0), statistics.sampleStandardDeviation, "Standard deviation should be 0 when only a single value was provided")
    assert.Equal(t, float64(0), statistics.coefficientOfVariation, "Coefficient of variation must be 0 when standard deviation is 0")
    assert.Equal(t, statistics.mean, statistics.sampleMean, "Sample mean must equal mean when all values are included in the sample set")
    assert.Equal(t, statistics.sum, statistics.sampleSum, "Sample sum must equal sum when all values are included in the sample set")
    assert.Equal(t, statistics.count, int64(statistics.sampleCount), "Sample count must count when all values are included in the sample set")
}

// Test that the Accumulator produces correct results when multiple values are received (and all included in the sample count) within the interval
func TestAccumulatorCycle_SomeValues(t *testing.T) {
    var targetSampleCount = 100 // the sample count is larger than the number of values provided within this test
    
    var accumulator = NewAccumulator(targetSampleCount)
    
    // this test covers a scenario in which only several values are received within the interval
    accumulator.Include(2)
    accumulator.Include(4)
    accumulator.Include(6)
    
    var expectedStandardDeviation = math.Sqrt(8)
    
    // finalise the accumulator and gather statistics
    var statistics = accumulator.Finalise()

    assert.Equal(t, int64(3), statistics.count, "Unexpected count")
    assert.Equal(t, float64(4), statistics.mean, "Unexpected mean")
    assert.Equal(t, float64(2), statistics.minimum, "Unexpected minimum")
    assert.Equal(t, float64(6), statistics.maximum, "Unexpected maximum")
    assert.Equal(t, float64(12), statistics.sum, "Unexpected sum")
    
    assert.Equal(t, expectedStandardDeviation, statistics.sampleStandardDeviation, "Unexpected standard deviation")
    assert.Equal(t, statistics.mean / expectedStandardDeviation, statistics.coefficientOfVariation, "Unexpected coefficient of variation")
    
    assert.Equal(t, statistics.mean, statistics.sampleMean, "Sample mean must equal mean when all values are included in the sample set")
    assert.Equal(t, statistics.sum, statistics.sampleSum, "Sample sum must equal sum when all values are included in the sample set")
    assert.Equal(t, statistics.count, int64(statistics.sampleCount), "Sample count must count when all values are included in the sample set")
}

// Test that the Accumulator produces correct results when multiple values (where some are excluded from the sample count) are received within the interval
func TestAccumulatorCycle_ManyValues(t *testing.T) {
    var targetSampleCount = 100 // the sample count is larger than the number of values provided within this test
    
    var accumulator = NewAccumulator(targetSampleCount)
    
    var overallSum = 0.0
    var overallCount = 0
    
    var values = [5]float64{2.0,4.0,6.0,8.0,10.0}
    var mean = 6.0
    
    var addedValues = make([]float64, 0, 100)
    
    var squareError = 0.0
    
    for _,v := range values {
        var valuesToAdd = targetSampleCount
        
        if v == 6 {
            valuesToAdd = 5
        }
        
        for i:=0;i<valuesToAdd;i++ {
            accumulator.Include(v)
            overallSum += v
            overallCount++
            squareError += math.Pow(v-mean,2)
            addedValues = append(addedValues, v)
        }    
    }
    
    //var standardDeviation = math.Sqrt(squareError)
    var overallMean = overallSum / float64(overallCount)
    
    // finalise the accumulator and gather statistics
    var statistics = accumulator.Finalise()

    assert.Equal(t, int64(overallCount), statistics.count, "Unexpected count")
    assert.Equal(t, overallMean, statistics.mean, "Unexpected mean")
    assert.Equal(t, 2.0, statistics.minimum, "Unexpected minimum")
    assert.Equal(t, 10.0, statistics.maximum, "Unexpected maximum")
    assert.Equal(t, overallSum, statistics.sum, "Unexpected sum")
    
    /*
    assert.Equal(t, standardDeviation, statistics.sampleStandardDeviation, "Unexpected standard deviation")
    assert.Equal(t, overallMean / standardDeviation, statistics.coefficientOfVariation, "Unexpected coefficient of variation")
    assert.Equal(t, float64(4), statistics.sampleMean, "Sample mean must equal mean when all values are included in the sample set")
    assert.Equal(t, statistics.sum, statistics.sampleSum, "Sample sum must equal sum when all values are included in the sample set")
    assert.Equal(t, statistics.count, int64(statistics.sampleCount), "Sample count must count when all values are included in the sample set")
    */
}
