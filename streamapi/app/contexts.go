//************************************************************************//
// API "Stream Statistics API": Application Contexts
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
	"golang.org/x/net/context"
	"strconv"
	"time"
)

// AddOrdinalValuesContext provides the OrdinalValues add action context.
type AddOrdinalValuesContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service *goa.Service
	Ordinal int
	Stream  string
	Value   float64
}

// NewAddOrdinalValuesContext parses the incoming request URL and body, performs validations and creates the
// context used by the OrdinalValues controller add action.
func NewAddOrdinalValuesContext(ctx context.Context, service *goa.Service) (*AddOrdinalValuesContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := AddOrdinalValuesContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramOrdinal := req.Params["ordinal"]
	if len(paramOrdinal) > 0 {
		rawOrdinal := paramOrdinal[0]
		if ordinal, err2 := strconv.Atoi(rawOrdinal); err2 == nil {
			rctx.Ordinal = ordinal
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("ordinal", rawOrdinal, "integer"))
		}
	}
	paramStream := req.Params["stream"]
	if len(paramStream) > 0 {
		rawStream := paramStream[0]
		rctx.Stream = rawStream
	}
	paramValue := req.Params["value"]
	if len(paramValue) > 0 {
		rawValue := paramValue[0]
		if value, err2 := strconv.ParseFloat(rawValue, 64); err2 == nil {
			rctx.Value = value
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("value", rawValue, "number"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *AddOrdinalValuesContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// PushOrdinalValuesContext provides the OrdinalValues push action context.
type PushOrdinalValuesContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service *goa.Service
	Payload *PushOrdinalValuesPayload
}

// NewPushOrdinalValuesContext parses the incoming request URL and body, performs validations and creates the
// context used by the OrdinalValues controller push action.
func NewPushOrdinalValuesContext(ctx context.Context, service *goa.Service) (*PushOrdinalValuesContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := PushOrdinalValuesContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	return &rctx, err
}

// pushOrdinalValuesPayload is the OrdinalValues push action payload.
type pushOrdinalValuesPayload struct {
	// The ordinal position within the stream
	Ordinal *int `json:"ordinal,omitempty" xml:"ordinal,omitempty"`
	// Identifies the stream that the ordinal value relates to
	Stream *string `json:"stream,omitempty" xml:"stream,omitempty"`
	// The value at the ordinal position
	Value *float64 `json:"value,omitempty" xml:"value,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *pushOrdinalValuesPayload) Validate() (err error) {
	if payload.Stream == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "stream"))
	}
	if payload.Ordinal == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "ordinal"))
	}
	if payload.Value == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "value"))
	}

	return
}

// Publicize creates PushOrdinalValuesPayload from pushOrdinalValuesPayload
func (payload *pushOrdinalValuesPayload) Publicize() *PushOrdinalValuesPayload {
	var pub PushOrdinalValuesPayload
	if payload.Ordinal != nil {
		pub.Ordinal = *payload.Ordinal
	}
	if payload.Stream != nil {
		pub.Stream = *payload.Stream
	}
	if payload.Value != nil {
		pub.Value = *payload.Value
	}
	return &pub
}

// PushOrdinalValuesPayload is the OrdinalValues push action payload.
type PushOrdinalValuesPayload struct {
	// The ordinal position within the stream
	Ordinal int `json:"ordinal" xml:"ordinal"`
	// Identifies the stream that the ordinal value relates to
	Stream string `json:"stream" xml:"stream"`
	// The value at the ordinal position
	Value float64 `json:"value" xml:"value"`
}

// Validate runs the validation rules defined in the design.
func (payload *PushOrdinalValuesPayload) Validate() (err error) {
	if payload.Stream == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "stream"))
	}

	return
}

// OK sends a HTTP response with status code 200.
func (ctx *PushOrdinalValuesContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// RegisterOrdinalValuesContext provides the OrdinalValues register action context.
type RegisterOrdinalValuesContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service *goa.Service
	Payload *RegisterOrdinalValuesPayload
}

// NewRegisterOrdinalValuesContext parses the incoming request URL and body, performs validations and creates the
// context used by the OrdinalValues controller register action.
func NewRegisterOrdinalValuesContext(ctx context.Context, service *goa.Service) (*RegisterOrdinalValuesContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := RegisterOrdinalValuesContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	return &rctx, err
}

// registerOrdinalValuesPayload is the OrdinalValues register action payload.
type registerOrdinalValuesPayload struct {
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

// Validate runs the validation rules defined in the design.
func (payload *registerOrdinalValuesPayload) Validate() (err error) {
	if payload.Stream == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "stream"))
	}
	if payload.IntervalSize == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "intervalSize"))
	}
	if payload.MaxIntervalLag == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "maxIntervalLag"))
	}
	if payload.TargetSampleSize == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "targetSampleSize"))
	}

	return
}

// Publicize creates RegisterOrdinalValuesPayload from registerOrdinalValuesPayload
func (payload *registerOrdinalValuesPayload) Publicize() *RegisterOrdinalValuesPayload {
	var pub RegisterOrdinalValuesPayload
	if payload.IntervalSize != nil {
		pub.IntervalSize = *payload.IntervalSize
	}
	if payload.MaxIntervalLag != nil {
		pub.MaxIntervalLag = *payload.MaxIntervalLag
	}
	if payload.Stream != nil {
		pub.Stream = *payload.Stream
	}
	if payload.Tags != nil {
		pub.Tags = payload.Tags
	}
	if payload.TargetSampleSize != nil {
		pub.TargetSampleSize = *payload.TargetSampleSize
	}
	return &pub
}

// RegisterOrdinalValuesPayload is the OrdinalValues register action payload.
type RegisterOrdinalValuesPayload struct {
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

// Validate runs the validation rules defined in the design.
func (payload *RegisterOrdinalValuesPayload) Validate() (err error) {
	if payload.Stream == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "stream"))
	}

	return
}

// OK sends a HTTP response with status code 200.
func (ctx *RegisterOrdinalValuesContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// StatisticsOrdinalValuesContext provides the OrdinalValues statistics action context.
type StatisticsOrdinalValuesContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service *goa.Service
	Payload *StatisticsOrdinalValuesPayload
}

// NewStatisticsOrdinalValuesContext parses the incoming request URL and body, performs validations and creates the
// context used by the OrdinalValues controller statistics action.
func NewStatisticsOrdinalValuesContext(ctx context.Context, service *goa.Service) (*StatisticsOrdinalValuesContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := StatisticsOrdinalValuesContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	return &rctx, err
}

// statisticsOrdinalValuesPayload is the OrdinalValues statistics action payload.
type statisticsOrdinalValuesPayload struct {
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

// Publicize creates StatisticsOrdinalValuesPayload from statisticsOrdinalValuesPayload
func (payload *statisticsOrdinalValuesPayload) Publicize() *StatisticsOrdinalValuesPayload {
	var pub StatisticsOrdinalValuesPayload
	if payload.MaxDateTime != nil {
		pub.MaxDateTime = payload.MaxDateTime
	}
	if payload.MaxOrdinal != nil {
		pub.MaxOrdinal = payload.MaxOrdinal
	}
	if payload.MergeIntervals != nil {
		pub.MergeIntervals = payload.MergeIntervals
	}
	if payload.MergeStreams != nil {
		pub.MergeStreams = payload.MergeStreams
	}
	if payload.MinDateTime != nil {
		pub.MinDateTime = payload.MinDateTime
	}
	if payload.MinOrdinal != nil {
		pub.MinOrdinal = payload.MinOrdinal
	}
	if payload.StreamMatchCriteria != nil {
		pub.StreamMatchCriteria = payload.StreamMatchCriteria.Publicize()
	}
	return &pub
}

// StatisticsOrdinalValuesPayload is the OrdinalValues statistics action payload.
type StatisticsOrdinalValuesPayload struct {
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

// OK sends a HTTP response with status code 200.
func (ctx *StatisticsOrdinalValuesContext) OK(r *GoaStatisticsresults) error {
	ctx.ResponseData.Header().Set("Content-Type", "vnd.application/goa.statisticsresults")
	return ctx.Service.Send(ctx.Context, 200, r)
}

// TagOrdinalValuesContext provides the OrdinalValues tag action context.
type TagOrdinalValuesContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service *goa.Service
	Payload *TagOrdinalValuesPayload
}

// NewTagOrdinalValuesContext parses the incoming request URL and body, performs validations and creates the
// context used by the OrdinalValues controller tag action.
func NewTagOrdinalValuesContext(ctx context.Context, service *goa.Service) (*TagOrdinalValuesContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := TagOrdinalValuesContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	return &rctx, err
}

// tagOrdinalValuesPayload is the OrdinalValues tag action payload.
type tagOrdinalValuesPayload struct {
	// If true, previously assigned tags will be cleared
	ClearAll *bool `json:"clearAll,omitempty" xml:"clearAll,omitempty"`
	// Identifies the stream that the definition relates to
	Stream *string `json:"stream,omitempty" xml:"stream,omitempty"`
	// An array of tags to be assigned
	TagsToAssign []string `json:"tagsToAssign,omitempty" xml:"tagsToAssign,omitempty"`
	// An array of tags to be unassigned
	TagsToUnassign []string `json:"tagsToUnassign,omitempty" xml:"tagsToUnassign,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *tagOrdinalValuesPayload) Validate() (err error) {
	if payload.Stream == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "stream"))
	}

	return
}

// Publicize creates TagOrdinalValuesPayload from tagOrdinalValuesPayload
func (payload *tagOrdinalValuesPayload) Publicize() *TagOrdinalValuesPayload {
	var pub TagOrdinalValuesPayload
	if payload.ClearAll != nil {
		pub.ClearAll = payload.ClearAll
	}
	if payload.Stream != nil {
		pub.Stream = *payload.Stream
	}
	if payload.TagsToAssign != nil {
		pub.TagsToAssign = payload.TagsToAssign
	}
	if payload.TagsToUnassign != nil {
		pub.TagsToUnassign = payload.TagsToUnassign
	}
	return &pub
}

// TagOrdinalValuesPayload is the OrdinalValues tag action payload.
type TagOrdinalValuesPayload struct {
	// If true, previously assigned tags will be cleared
	ClearAll *bool `json:"clearAll,omitempty" xml:"clearAll,omitempty"`
	// Identifies the stream that the definition relates to
	Stream string `json:"stream" xml:"stream"`
	// An array of tags to be assigned
	TagsToAssign []string `json:"tagsToAssign,omitempty" xml:"tagsToAssign,omitempty"`
	// An array of tags to be unassigned
	TagsToUnassign []string `json:"tagsToUnassign,omitempty" xml:"tagsToUnassign,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *TagOrdinalValuesPayload) Validate() (err error) {
	if payload.Stream == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "stream"))
	}

	return
}

// OK sends a HTTP response with status code 200.
func (ctx *TagOrdinalValuesContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}
