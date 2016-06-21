package stream

import "testing"
import "github.com/stretchr/testify/assert"
import "math"

// Test that the Accumulator produces correct results when no values are received within the interval
func TestAccumulatorCycle_NoValue(t *testing.T) {
    var targetSampleCount = uint32(100) // the sample count is larger than the number of values provided within this test
    
    var accumulator = NewAccumulator("stream", 0, math.MaxInt64, OrdinalInterval, targetSampleCount)
    
    // this test covers a scenario in which no values are received within the interval
    // .. no values are provided to the accumulator ..
        
    // finalise the accumulator and gather statistics
    var statistics = accumulator.Finalise()
    
    assert.Equal(t, uint64(0), statistics.Count, "Count should be 0 when no values where provided to the accumulator")
    assert.Equal(t, statistics.Count, uint64(statistics.SampleCount), "Sample count must count when all values are included in the sample set")
}

// Test that the Accumulator produces correct results when only a single value is received within the interval
func TestAccumulatorCycle_SingleValue(t *testing.T) {
    var targetSampleCount = uint32(100) // the sample count is larger than the number of values provided within this test
    
    var accumulator = NewAccumulator("stream", 0, math.MaxInt64, OrdinalInterval, targetSampleCount)
    
    var value float64 = 50
    
    // this test covers a scenario in which only 1 value is received within the interval
    accumulator.Include(NewOrdinalValue("test",1,float64(value)))
    
    // finalise the accumulator and gather statistics
    var statistics = accumulator.Finalise()
    
    assert.Equal(t, uint64(1), statistics.Count, "Count should be 1 when a single value was provided to the accumulator")
    assert.Equal(t, value, statistics.Mean, "Mean should equal the single value when only a single value was provided to the accumulator")
    assert.Equal(t, value, statistics.Minimum, "Minimum should equal the single value when only a single value was provided to the accumulator")
    assert.Equal(t, value, statistics.Maximum, "Maximum should equal the single value when only a single value was provided to the accumulator")
    assert.Equal(t, value, statistics.Sum, "Sum should equal the single value when only a single value was provided to the accumulator")
    assert.Equal(t, float64(0), statistics.SampleStandardDeviation, "Standard deviation should be 0 when only a single value was provided")
    assert.Equal(t, float64(0), statistics.CoefficientOfVariation, "Coefficient of variation must be 0 when standard deviation is 0")
    assert.Equal(t, statistics.Mean, statistics.SampleMean, "Sample mean must equal mean when all values are included in the sample set")
    assert.Equal(t, statistics.Sum, statistics.SampleSum, "Sample sum must equal sum when all values are included in the sample set")
    assert.Equal(t, statistics.Count, uint64(statistics.SampleCount), "Sample count must count when all values are included in the sample set")
}

// Test that the Accumulator produces correct results when multiple values are received (and all included in the sample count) within the interval
func TestAccumulatorCycle_SomeValues(t *testing.T) {
    var targetSampleCount = uint32(100) // the sample count is larger than the number of values provided within this test
    
    var accumulator = NewAccumulator("stream", 0, math.MaxInt64, OrdinalInterval, targetSampleCount)
    
    // this test covers a scenario in which only several values are received within the interval
    accumulator.Include(NewOrdinalValue("test",1,2))
    accumulator.Include(NewOrdinalValue("test",1,4))
    accumulator.Include(NewOrdinalValue("test",1,6))
    
    var expectedStandardDeviation = math.Sqrt(8.0/3.0)
    
    // finalise the accumulator and gather statistics
    var statistics = accumulator.Finalise()

    assert.Equal(t, uint64(3), statistics.Count, "Unexpected count")
    assert.Equal(t, float64(4), statistics.Mean, "Unexpected mean")
    assert.Equal(t, float64(2), statistics.Minimum, "Unexpected minimum")
    assert.Equal(t, float64(6), statistics.Maximum, "Unexpected maximum")
    assert.Equal(t, float64(12), statistics.Sum, "Unexpected sum")
    
    assert.Equal(t, expectedStandardDeviation, statistics.SampleStandardDeviation, "Unexpected standard deviation")
    assert.Equal(t, statistics.Mean / expectedStandardDeviation, statistics.CoefficientOfVariation, "Unexpected coefficient of variation")
    
    assert.Equal(t, statistics.Mean, statistics.SampleMean, "Sample mean must equal mean when all values are included in the sample set")
    assert.Equal(t, statistics.Sum, statistics.SampleSum, "Sample sum must equal sum when all values are included in the sample set")
    assert.Equal(t, statistics.Count, uint64(statistics.SampleCount), "Sample count must count when all values are included in the sample set")
}

// Test that the Accumulator produces correct results when multiple values (where some are excluded from the sample count) are received within the interval
func TestAccumulatorCycle_ManyValues(t *testing.T) {
    var targetSampleCount = uint32(100) // the sample count is larger than the number of values provided within this test
    
    var accumulator = NewAccumulator("stream", 0, math.MaxInt64, OrdinalInterval, targetSampleCount)
    
    var overallSum = 0.0
    var overallCount = 0
    
    var values = [5]float64{2.0,4.0,6.0,8.0,10.0}
    var addedValues = make([]float64, 0, 100)
    
    for _,v := range values {
        var valuesToAdd = targetSampleCount
        
        if v == 6 {
            valuesToAdd = 5
        }
        
        for i:=uint32(0);i<valuesToAdd;i++ {
            accumulator.Include(NewOrdinalValue("test",1,v))
            overallSum += v
            overallCount++
            addedValues = append(addedValues, v)
        }    
    }
    
    var overallMean = overallSum / float64(overallCount)
    
    // finalise the accumulator and gather statistics
    var statistics = accumulator.Finalise()

    assert.Equal(t, uint64(overallCount), statistics.Count, "Unexpected count")
    assert.Equal(t, overallMean, statistics.Mean, "Unexpected mean")
    assert.Equal(t, 2.0, statistics.Minimum, "Unexpected minimum")
    assert.Equal(t, 10.0, statistics.Maximum, "Unexpected maximum")
    assert.Equal(t, overallSum, statistics.Sum, "Unexpected sum")
    
    // determine the sample set that would have been used
    var samplingRateDenominator = 1
    var maxSamplingRate = 1.5
    var expectedSampleCount = overallCount
    var rate = float64(expectedSampleCount) / float64(targetSampleCount)
    
    for rate > maxSamplingRate {
        samplingRateDenominator *= 2
        expectedSampleCount = int(math.Floor(float64(expectedSampleCount) / 2.0))
        rate = float64(expectedSampleCount) / float64(targetSampleCount)
    }
    
    var sampleValues = make([]float64, 0, 100)
    var sampleCount = uint32(0)
    var sampleSum = 0.0
    
    for k,v := range addedValues {
        if (k + 1) % samplingRateDenominator == 0 {
            sampleValues = append(sampleValues, v)
            sampleCount++
            sampleSum += v
        }
    }
    
    var sampleMean = sampleSum / float64(sampleCount)
    var sumSqSampleErr = 0.0
    
    for _,v := range sampleValues {
        sumSqSampleErr += math.Pow(v-sampleMean,2)
    }
    
    var sampleStandardDeviation = math.Sqrt(sumSqSampleErr / float64(sampleCount))
    
    assert.Equal(t, sampleStandardDeviation, statistics.SampleStandardDeviation, "Unexpected standard deviation")
    assert.Equal(t, sampleMean / sampleStandardDeviation, statistics.CoefficientOfVariation, "Unexpected coefficient of variation")
    assert.Equal(t, sampleMean, statistics.SampleMean, "Sample mean must equal mean when all values are included in the sample set")
    assert.Equal(t, sampleSum, statistics.SampleSum, "Sample sum must equal sum when all values are included in the sample set")
    assert.Equal(t, sampleCount, statistics.SampleCount, "Sample count must count when all values are included in the sample set")
}

// Test that the Accumulator produces correct results when multiple values (where some are excluded from the sample count) are received within the interval
func TestAccumulatorCycle_ManyValuesWithChannels(t *testing.T) {
    var targetSampleCount = uint32(100) // the sample count is larger than the number of values provided within this test
    
    var accumulator = NewAccumulator("stream", 0, math.MaxInt64, OrdinalInterval, targetSampleCount)
    
    input := make(chan OrdinalValue)
    output := make(chan IntervalStatistics)
    doneChannel := make(chan bool)
    
    go accumulator.Accumulate(input, output, doneChannel)
    
    var overallSum = 0.0
    var overallCount = 0
    
    var values = [5]float64{2.0,4.0,6.0,8.0,10.0}
    var addedValues = make([]float64, 0, 100)
    
    for _,v := range values {
        var valuesToAdd = targetSampleCount
        
        if v == 6 {
            valuesToAdd = 5
        }
        
        for i:=uint32(0);i<valuesToAdd;i++ {
            input <- NewOrdinalValue("test",1,v)
            overallSum += v
            overallCount++
            addedValues = append(addedValues, v)
        }    
    }
    close(input)
    
    var overallMean = overallSum / float64(overallCount)
    
    // finalise the accumulator and gather statistics
    statistics := <- output
    <- doneChannel

    assert.Equal(t, uint64(overallCount), statistics.Count, "Unexpected count")
    assert.Equal(t, overallMean, statistics.Mean, "Unexpected mean")
    assert.Equal(t, 2.0, statistics.Minimum, "Unexpected minimum")
    assert.Equal(t, 10.0, statistics.Maximum, "Unexpected maximum")
    assert.Equal(t, overallSum, statistics.Sum, "Unexpected sum")
    
    // determine the sample set that would have been used
    var samplingRateDenominator = 1
    var maxSamplingRate = 1.5
    var expectedSampleCount = overallCount
    var rate = float64(expectedSampleCount) / float64(targetSampleCount)
    
    for rate > maxSamplingRate {
        samplingRateDenominator *= 2
        expectedSampleCount = int(math.Floor(float64(expectedSampleCount) / 2.0))
        rate = float64(expectedSampleCount) / float64(targetSampleCount)
    }
    
    var sampleValues = make([]float64, 0, 100)
    var sampleCount = uint32(0)
    var sampleSum = 0.0
    
    for k,v := range addedValues {
        if (k + 1) % samplingRateDenominator == 0 {
            sampleValues = append(sampleValues, v)
            sampleCount++
            sampleSum += v
        }
    }
    
    var sampleMean = sampleSum / float64(sampleCount)
    var sumSqSampleErr = 0.0
    
    for _,v := range sampleValues {
        sumSqSampleErr += math.Pow(v-sampleMean,2)
    }
    
    var sampleStandardDeviation = math.Sqrt(sumSqSampleErr / float64(sampleCount))
    
    assert.Equal(t, sampleStandardDeviation, statistics.SampleStandardDeviation, "Unexpected standard deviation")
    assert.Equal(t, sampleMean / sampleStandardDeviation, statistics.CoefficientOfVariation, "Unexpected coefficient of variation")
    assert.Equal(t, sampleMean, statistics.SampleMean, "Sample mean must equal mean when all values are included in the sample set")
    assert.Equal(t, sampleSum, statistics.SampleSum, "Sample sum must equal sum when all values are included in the sample set")
    assert.Equal(t, sampleCount, statistics.SampleCount, "Sample count must count when all values are included in the sample set")
}
