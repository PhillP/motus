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
    StreamKey                   string
    IntervalStart               int64
    IntervalEnd                 int64
    IntervalType                IntervalType
    Minimum                     float64
    Maximum                     float64
    Mean                        float64
    Count                       uint64
    Sum                         float64
    SampleMean                  float64
    SampleSum                   float64
    SampleCount                 uint32
    SampleStandardDeviation     float64
    CoefficientOfVariation      float64
}
