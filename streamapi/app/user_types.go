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

import (
	"github.com/goadesign/goa"
	"time"
)

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

// statisticsSearchCriteria user type.
type statisticsSearchCriteria struct {
	// Specifies a maximum date time used to restrict the interval statistics returned.  Only statistics for intervals that are for a time range up until this date time value will be returned.
	MaxDateTime *time.Time `json:"maxDateTime,omitempty" xml:"maxDateTime,omitempty"`
	// Specifies a maximum ordinal value used to restrict the interval statistics returned.  Only statistics for intervals that end on or before this ordinal value will be returned.
	MaxOrdinal *int `json:"maxOrdinal,omitempty" xml:"maxOrdinal,omitempty"`
	// If true, results across multiple intervals will be merged together to produce a summary result.
	MergeIntervals *bool `json:"mergeIntervals,omitempty" xml:"mergeIntervals,omitempty"`
	// If true, results from multiple streams will be merged together to produce a summary result.
	MergeStreams *bool `json:"mergeStreams,omitempty" xml:"mergeStreams,omitempty"`
	// Specifies a minimum date time used to restrict the interval statistics returned.  Only statistics for intervals that are for a time range on or after this date time value will be returned.
	MinDateTime *time.Time `json:"minDateTime,omitempty" xml:"minDateTime,omitempty"`
	// Specifies a minimum ordinal value used to restrict the interval statistics returned.  Only statistics for intervals that begin on or after this ordinal value will be returned.
	MinOrdinal *int `json:"minOrdinal,omitempty" xml:"minOrdinal,omitempty"`
	// Specifies the criteria by which streams are to be matched
	StreamMatchCriteria *streamMatchCriteria `json:"streamMatchCriteria,omitempty" xml:"streamMatchCriteria,omitempty"`
}

// Publicize creates StatisticsSearchCriteria from statisticsSearchCriteria
func (ut *statisticsSearchCriteria) Publicize() *StatisticsSearchCriteria {
	var pub StatisticsSearchCriteria
	if ut.MaxDateTime != nil {
		pub.MaxDateTime = ut.MaxDateTime
	}
	if ut.MaxOrdinal != nil {
		pub.MaxOrdinal = ut.MaxOrdinal
	}
	if ut.MergeIntervals != nil {
		pub.MergeIntervals = ut.MergeIntervals
	}
	if ut.MergeStreams != nil {
		pub.MergeStreams = ut.MergeStreams
	}
	if ut.MinDateTime != nil {
		pub.MinDateTime = ut.MinDateTime
	}
	if ut.MinOrdinal != nil {
		pub.MinOrdinal = ut.MinOrdinal
	}
	if ut.StreamMatchCriteria != nil {
		pub.StreamMatchCriteria = ut.StreamMatchCriteria.Publicize()
	}
	return &pub
}

// StatisticsSearchCriteria user type.
type StatisticsSearchCriteria struct {
	// Specifies a maximum date time used to restrict the interval statistics returned.  Only statistics for intervals that are for a time range up until this date time value will be returned.
	MaxDateTime *time.Time `json:"maxDateTime,omitempty" xml:"maxDateTime,omitempty"`
	// Specifies a maximum ordinal value used to restrict the interval statistics returned.  Only statistics for intervals that end on or before this ordinal value will be returned.
	MaxOrdinal *int `json:"maxOrdinal,omitempty" xml:"maxOrdinal,omitempty"`
	// If true, results across multiple intervals will be merged together to produce a summary result.
	MergeIntervals *bool `json:"mergeIntervals,omitempty" xml:"mergeIntervals,omitempty"`
	// If true, results from multiple streams will be merged together to produce a summary result.
	MergeStreams *bool `json:"mergeStreams,omitempty" xml:"mergeStreams,omitempty"`
	// Specifies a minimum date time used to restrict the interval statistics returned.  Only statistics for intervals that are for a time range on or after this date time value will be returned.
	MinDateTime *time.Time `json:"minDateTime,omitempty" xml:"minDateTime,omitempty"`
	// Specifies a minimum ordinal value used to restrict the interval statistics returned.  Only statistics for intervals that begin on or after this ordinal value will be returned.
	MinOrdinal *int `json:"minOrdinal,omitempty" xml:"minOrdinal,omitempty"`
	// Specifies the criteria by which streams are to be matched
	StreamMatchCriteria *StreamMatchCriteria `json:"streamMatchCriteria,omitempty" xml:"streamMatchCriteria,omitempty"`
}

// streamDefinition user type.
type streamDefinition struct {
	// The ordinal position within the stream
	IntervalSize *int `json:"intervalSize,omitempty" xml:"intervalSize,omitempty"`
	// The value at the ordinal position
	MaxIntervalLag *int `json:"maxIntervalLag,omitempty" xml:"maxIntervalLag,omitempty"`
	// Identifies the stream that the definition relates to
	Stream *string `json:"stream,omitempty" xml:"stream,omitempty"`
	// A set of tag values to be assigned to the stream
	Tags []string `json:"tags,omitempty" xml:"tags,omitempty"`
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
	if ut.Tags != nil {
		pub.Tags = ut.Tags
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
	// A set of tag values to be assigned to the stream
	Tags []string `json:"tags,omitempty" xml:"tags,omitempty"`
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

// streamMatchCriteria user type.
type streamMatchCriteria struct {
	// An optional array of tags.  Streams tagged with any of these tags will be excluded
	ExcludeWithAnyTags []string `json:"excludeWithAnyTags,omitempty" xml:"excludeWithAnyTags,omitempty"`
	// An optional array of tags. Streams tagged with all of these tags will be included
	IncludeWithAllTags []string `json:"includeWithAllTags,omitempty" xml:"includeWithAllTags,omitempty"`
	// An optional array of streamKeys used to select streams
	StreamKeys []string `json:"streamKeys,omitempty" xml:"streamKeys,omitempty"`
}

// Publicize creates StreamMatchCriteria from streamMatchCriteria
func (ut *streamMatchCriteria) Publicize() *StreamMatchCriteria {
	var pub StreamMatchCriteria
	if ut.ExcludeWithAnyTags != nil {
		pub.ExcludeWithAnyTags = ut.ExcludeWithAnyTags
	}
	if ut.IncludeWithAllTags != nil {
		pub.IncludeWithAllTags = ut.IncludeWithAllTags
	}
	if ut.StreamKeys != nil {
		pub.StreamKeys = ut.StreamKeys
	}
	return &pub
}

// StreamMatchCriteria user type.
type StreamMatchCriteria struct {
	// An optional array of tags.  Streams tagged with any of these tags will be excluded
	ExcludeWithAnyTags []string `json:"excludeWithAnyTags,omitempty" xml:"excludeWithAnyTags,omitempty"`
	// An optional array of tags. Streams tagged with all of these tags will be included
	IncludeWithAllTags []string `json:"includeWithAllTags,omitempty" xml:"includeWithAllTags,omitempty"`
	// An optional array of streamKeys used to select streams
	StreamKeys []string `json:"streamKeys,omitempty" xml:"streamKeys,omitempty"`
}

// tagAssignmentDefinition user type.
type tagAssignmentDefinition struct {
	// If true, previously assigned tags will be cleared
	ClearAll *bool `json:"clearAll,omitempty" xml:"clearAll,omitempty"`
	// Identifies the stream that the definition relates to
	Stream *string `json:"stream,omitempty" xml:"stream,omitempty"`
	// An array of tags to be assigned
	TagsToAssign []string `json:"tagsToAssign,omitempty" xml:"tagsToAssign,omitempty"`
	// An array of tags to be unassigned
	TagsToUnassign []string `json:"tagsToUnassign,omitempty" xml:"tagsToUnassign,omitempty"`
}

// Validate validates the tagAssignmentDefinition type instance.
func (ut *tagAssignmentDefinition) Validate() (err error) {
	if ut.Stream == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "stream"))
	}

	return
}

// Publicize creates TagAssignmentDefinition from tagAssignmentDefinition
func (ut *tagAssignmentDefinition) Publicize() *TagAssignmentDefinition {
	var pub TagAssignmentDefinition
	if ut.ClearAll != nil {
		pub.ClearAll = ut.ClearAll
	}
	if ut.Stream != nil {
		pub.Stream = *ut.Stream
	}
	if ut.TagsToAssign != nil {
		pub.TagsToAssign = ut.TagsToAssign
	}
	if ut.TagsToUnassign != nil {
		pub.TagsToUnassign = ut.TagsToUnassign
	}
	return &pub
}

// TagAssignmentDefinition user type.
type TagAssignmentDefinition struct {
	// If true, previously assigned tags will be cleared
	ClearAll *bool `json:"clearAll,omitempty" xml:"clearAll,omitempty"`
	// Identifies the stream that the definition relates to
	Stream string `json:"stream" xml:"stream"`
	// An array of tags to be assigned
	TagsToAssign []string `json:"tagsToAssign,omitempty" xml:"tagsToAssign,omitempty"`
	// An array of tags to be unassigned
	TagsToUnassign []string `json:"tagsToUnassign,omitempty" xml:"tagsToUnassign,omitempty"`
}

// Validate validates the TagAssignmentDefinition type instance.
func (ut *TagAssignmentDefinition) Validate() (err error) {
	if ut.Stream == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "stream"))
	}

	return
}
