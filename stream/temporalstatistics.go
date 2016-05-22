package stream

import "time"

// TemporalStatistics is a set of statistics generate for a period of time based on processed events
type TemporalStatistics struct {
    fromTime                    time.Time
    untilTime                   time.Time
    hasValue                    bool
    minimum                     float64
    maximum                     float64
    mean                        float64
    count                       uint64
    standardDeviation           float64
    coefficientOfVariation      float64
}