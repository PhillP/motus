//************************************************************************//
// API "Stream Statistics API": Application User Types
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --design=github.com/phillp/motus/apidesign/stream
// --out=streamapi
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "github.com/goadesign/goa"

// ordinalValue user type.
type ordinalValue struct {
	// The ordinal position within the stream
	Ordinal *int `json:"ordinal,omitempty" xml:"ordinal,omitempty"`
	// Identifies the stream that the ordinal value relates to
	Stream *string `json:"stream,omitempty" xml:"stream,omitempty"`
	// The value at the ordinal position
	Value *float64 `json:"value,omitempty" xml:"value,omitempty"`
}

// Validate validates the ordinalValue type instance.
func (ut *ordinalValue) Validate() (err error) {
	if ut.Stream == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "stream"))
	}
	if ut.Ordinal == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "ordinal"))
	}
	if ut.Value == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "value"))
	}

	return
}

// Publicize creates OrdinalValue from ordinalValue
func (ut *ordinalValue) Publicize() *OrdinalValue {
	var pub OrdinalValue
	if ut.Ordinal != nil {
		pub.Ordinal = *ut.Ordinal
	}
	if ut.Stream != nil {
		pub.Stream = *ut.Stream
	}
	if ut.Value != nil {
		pub.Value = *ut.Value
	}
	return &pub
}

// OrdinalValue user type.
type OrdinalValue struct {
	// The ordinal position within the stream
	Ordinal int `json:"ordinal" xml:"ordinal"`
	// Identifies the stream that the ordinal value relates to
	Stream string `json:"stream" xml:"stream"`
	// The value at the ordinal position
	Value float64 `json:"value" xml:"value"`
}

// Validate validates the OrdinalValue type instance.
func (ut *OrdinalValue) Validate() (err error) {
	if ut.Stream == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "stream"))
	}

	return
}

// streamDefinition user type.
type streamDefinition struct {
	// The ordinal position within the stream
	IntervalSize *int `json:"intervalSize,omitempty" xml:"intervalSize,omitempty"`
	// The value at the ordinal position
	MaxIntervalLag *int `json:"maxIntervalLag,omitempty" xml:"maxIntervalLag,omitempty"`
	// Identifies the stream that the definition relates to
	Stream *string `json:"stream,omitempty" xml:"stream,omitempty"`
	// The value at the ordinal position
	TargetSampleSize *int `json:"targetSampleSize,omitempty" xml:"targetSampleSize,omitempty"`
}

// Validate validates the streamDefinition type instance.
func (ut *streamDefinition) Validate() (err error) {
	if ut.Stream == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "stream"))
	}
	if ut.IntervalSize == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "intervalSize"))
	}
	if ut.MaxIntervalLag == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "maxIntervalLag"))
	}
	if ut.TargetSampleSize == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "targetSampleSize"))
	}

	return
}

// Publicize creates StreamDefinition from streamDefinition
func (ut *streamDefinition) Publicize() *StreamDefinition {
	var pub StreamDefinition
	if ut.IntervalSize != nil {
		pub.IntervalSize = *ut.IntervalSize
	}
	if ut.MaxIntervalLag != nil {
		pub.MaxIntervalLag = *ut.MaxIntervalLag
	}
	if ut.Stream != nil {
		pub.Stream = *ut.Stream
	}
	if ut.TargetSampleSize != nil {
		pub.TargetSampleSize = *ut.TargetSampleSize
	}
	return &pub
}

// StreamDefinition user type.
type StreamDefinition struct {
	// The ordinal position within the stream
	IntervalSize int `json:"intervalSize" xml:"intervalSize"`
	// The value at the ordinal position
	MaxIntervalLag int `json:"maxIntervalLag" xml:"maxIntervalLag"`
	// Identifies the stream that the definition relates to
	Stream string `json:"stream" xml:"stream"`
	// The value at the ordinal position
	TargetSampleSize int `json:"targetSampleSize" xml:"targetSampleSize"`
}

// Validate validates the StreamDefinition type instance.
func (ut *StreamDefinition) Validate() (err error) {
	if ut.Stream == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "stream"))
	}

	return
}
