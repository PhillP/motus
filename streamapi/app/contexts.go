//************************************************************************//
// API "Stream Statistics API": Application Contexts
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --design=github.com/phillp/motus/apidesign/stream
// --out=github.com/phillp/motus/streamapi
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"strconv"
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
