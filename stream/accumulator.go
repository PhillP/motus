package tempo

// Accumulator used to calculate statistics for a period of time by processing provided values
type Accumulator struct {
    hasValue                    bool
    minimum                     float64
    maximum                     float64
    count                       uint64
    sum                         float64
}

// NewAccumulator creates an accumulator
func NewAccumulator() (*Accumulator) {
    accumulator := Accumulator{}
    return &accumulator
}

// Include a new value within the accumulation
func (accumulator *Accumulator) Include(value float64) {
    
    if !accumulator.hasValue || accumulator.minimum > value {
        accumulator.minimum = value
    }
    
    if !accumulator.hasValue || accumulator.maximum < value {
        accumulator.maximum = value
    }
    
    if !accumulator.hasValue {
        accumulator.hasValue = true
    }
    
    accumulator.count++
    accumulator.sum = accumulator.sum + value
}