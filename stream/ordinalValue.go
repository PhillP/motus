package stream

import "time"

// OrdinalValue represents a value with an ordinal position within a stream
type OrdinalValue struct {
    StreamKey       string
    Ordinal         int64
    Value           float64
}

// NewOrdinalValue creates a new ordinal value
func NewOrdinalValue(streamKey string, ordinal int64, value float64) (OrdinalValue) {
    ordinalValue := OrdinalValue{ 
        StreamKey: streamKey,
        Ordinal: ordinal,
        Value: value}
        
    return ordinalValue
}

// NewOrdinalValueForTime creates a new ordinal value for a time
func NewOrdinalValueForTime(streamKey string, t time.Time, value float64) (OrdinalValue) {
    return NewOrdinalValue(streamKey, t.UnixNano(), value)
}