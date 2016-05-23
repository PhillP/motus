package stream

import "math"

// Accumulator used to calculate statistics for a period of time by processing provided values
type Accumulator struct {
    minimum                     float64
    maximum                     float64
    count                       int64
    sum                         float64
    sampleCount                 int
    targetSampleCount           int
    samplingRateDenominator     int
    sampleValues                []float64
    finalised                   bool
}


// NewAccumulator creates an accumulator
func NewAccumulator(targetSampleCount int) (*Accumulator) {
    accumulator := Accumulator{ 
        samplingRateDenominator: 1, 
        targetSampleCount: targetSampleCount,    
        sampleValues: make([]float64, 0, 100),
        finalised: false}
        
    return &accumulator
}

// Finalise calculates statistics from the accumulator and prevents any further accumulation
func (accumulator *Accumulator) Finalise() IntervalStatistics {
    
    // generate statistics based on the captured sample values
    
    var sampleSum float64
    var sampleMean float64
    var sampleStandardDeviation float64
    var coefficientOfVariation float64
    var mean float64
    
    if accumulator.count > 0 {
        mean = accumulator.sum / float64(accumulator.count)
    }
    
    if accumulator.sampleCount > 0 {
        // calculate sample mean
        for _,v := range accumulator.sampleValues {
            sampleSum += v    
        }
        sampleMean = sampleSum / float64(accumulator.sampleCount) 
    }
    
    var sumSquareError float64
    for _,v := range accumulator.sampleValues {
        sumSquareError += math.Pow(v - sampleMean, 2)    
    }
    sampleStandardDeviation = math.Sqrt(sumSquareError)
    
    if sampleStandardDeviation > 0 {
        coefficientOfVariation = sampleMean / sampleStandardDeviation
    }
    
    return IntervalStatistics{
        minimum: accumulator.minimum,
        maximum: accumulator.maximum,
        count: accumulator.count,
        sum: accumulator.sum,
        mean: mean,
        sampleCount: accumulator.sampleCount,
        sampleMean: sampleMean,
        sampleStandardDeviation: sampleStandardDeviation,
        sampleSum: sampleSum,
        coefficientOfVariation: coefficientOfVariation}
}

// Include a new value within the accumulation
func (accumulator *Accumulator) Include(value float64) {
    
    if accumulator.finalised {
        panic("Accumulator cannot include any more values after finalisation")
    }
    
    if accumulator.count == 0 || accumulator.minimum > value {
        accumulator.minimum = value
    }
    
    if accumulator.count == 0 || accumulator.maximum < value {
        accumulator.maximum = value
    }
    
    accumulator.count++
    accumulator.sum += value
    
    if accumulator.count % int64(accumulator.samplingRateDenominator) == 0 {
        accumulator.sampleValues = append(accumulator.sampleValues, value)
        accumulator.sampleCount++
        
        // check if there are now too many samples
        if float64(accumulator.sampleCount) > float64(accumulator.targetSampleCount) * .5 {
            // adjust the sampling rate
            accumulator.samplingRateDenominator *= 2
            
            var sampleSubset = make([]float64, 0)
            
            // and remove half the samples
            for i,v := range accumulator.sampleValues {
                if i % 2 == 0 {
                    sampleSubset = append(sampleSubset, v)        
                }
            }
            
            accumulator.sampleValues = sampleSubset
            accumulator.sampleCount = len(accumulator.sampleValues)
        }
    }
}