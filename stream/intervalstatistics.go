package stream

// IntervalType is a specifier used to differentiate between Ordinal and Time intervals
type IntervalType int

const (
    //OrdinalInterval indicates that intervals are determined based on ordinal position
    OrdinalInterval IntervalType = iota
    //TimeInterval indicates that intervals are determined based on time
    TimeInterval
)

// IntervalStatistics is a set of statistics generated based on processed events within a period of time or over a range or ordinal positions 
type IntervalStatistics struct {
    intervalStart               int64
    intervalEnd                 int64
    intervalType                IntervalType
    minimum                     float64
    maximum                     float64
    mean                        float64
    count                       uint64
    sum                         float64
    sampleMean                  float64
    sampleSum                   float64
    sampleCount                 uint32
    sampleStandardDeviation     float64
    coefficientOfVariation      float64
}
