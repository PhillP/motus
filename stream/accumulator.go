package stream

import "math"

// Accumulator used to calculate statistics for a period of time by processing provided values
type Accumulator struct {
    intervalStart               int64
    intervalEnd                 int64
    intervalType                IntervalType
    minimum                     float64
    maximum                     float64
    count                       uint64
    sum                         float64
    sampleCount                 uint32
    targetSampleCount           uint32
    samplingRateDenominator     uint32
    sampleValues                []OrdinalValue
    finalised                   bool
}

// NewAccumulator creates an accumulator
func NewAccumulator(intervalStart int64, intervalEnd int64, intervalType IntervalType, targetSampleCount uint32) (*Accumulator) {
    accumulator := Accumulator{ 
        intervalStart: intervalStart,
        intervalEnd: intervalEnd,
        intervalType: intervalType,
        samplingRateDenominator: 1, 
        targetSampleCount: targetSampleCount,    
        sampleValues: make([]OrdinalValue, 0, 100),
        finalised: false}
        
    return &accumulator
}

// Accumulate values from a channel
func (accumulator *Accumulator) Accumulate(input chan OrdinalValue, output chan IntervalStatistics, done chan bool) {
    for v := range input {
        accumulator.Include(v)
    }
    
    output <- accumulator.Finalise()
    done <- true
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
            sampleSum += v.Value    
        }
        sampleMean = sampleSum / float64(accumulator.sampleCount) 
    }
    
    var sumSquareError float64
    for _,v := range accumulator.sampleValues {
        sumSquareError += math.Pow(v.Value - sampleMean, 2)    
    }
    sampleStandardDeviation = math.Sqrt(sumSquareError / float64(accumulator.sampleCount))
    
    if sampleStandardDeviation > 0 {
        coefficientOfVariation = sampleMean / sampleStandardDeviation
    }
    
    return IntervalStatistics{
        IntervalStart: accumulator.intervalStart,
        IntervalEnd: accumulator.intervalEnd,
        IntervalType: accumulator.intervalType,
        Minimum: accumulator.minimum,
        Maximum: accumulator.maximum,
        Count: accumulator.count,
        Sum: accumulator.sum,
        Mean: mean,
        SampleCount: accumulator.sampleCount,
        SampleMean: sampleMean,
        SampleStandardDeviation: sampleStandardDeviation,
        SampleSum: sampleSum,
        CoefficientOfVariation: coefficientOfVariation}
}

// Include a new value within the accumulation
func (accumulator *Accumulator) Include(ordinalValue OrdinalValue) {
    
    value:=ordinalValue.Value
    
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
    
    if accumulator.count % uint64(accumulator.samplingRateDenominator) == 0 {
        accumulator.sampleValues = append(accumulator.sampleValues, ordinalValue)
        accumulator.sampleCount++
        
        // check if there are now too many samples
        if float64(accumulator.sampleCount) > float64(accumulator.targetSampleCount) * 1.5 {
            // adjust the sampling rate
            accumulator.samplingRateDenominator *= 2
            
            var sampleSubset = make([]OrdinalValue, 0)
            
            // and remove half the samples
            for i,v := range accumulator.sampleValues {
                if (i + 1) % 2 == 0 {
                    sampleSubset = append(sampleSubset, v)        
                }
            }
            
            accumulator.sampleValues = sampleSubset
            accumulator.sampleCount = uint32(len(accumulator.sampleValues))
        }
    }
}