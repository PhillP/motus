//************************************************************************//
// API "Stream Statistics API": Application Controllers
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
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// OrdinalValuesController is the controller interface for the OrdinalValues actions.
type OrdinalValuesController interface {
	goa.Muxer
	Add(*AddOrdinalValuesContext) error
	Push(*PushOrdinalValuesContext) error
	Register(*RegisterOrdinalValuesContext) error
	Statistics(*StatisticsOrdinalValuesContext) error
	Tag(*TagOrdinalValuesContext) error
}

// MountOrdinalValuesController "mounts" a OrdinalValues resource controller on the given service.
func MountOrdinalValuesController(service *goa.Service, ctrl OrdinalValuesController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewAddOrdinalValuesContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Add(rctx)
	}
	service.Mux.Handle("GET", "/api/add/:stream/:ordinal/:value", ctrl.MuxHandler("Add", h, nil))
	service.LogInfo("mount", "ctrl", "OrdinalValues", "action", "Add", "route", "GET /api/add/:stream/:ordinal/:value")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewPushOrdinalValuesContext(ctx, service)
		if err != nil {
			return err
		}
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*PushOrdinalValuesPayload)
		} else {
			return goa.ErrInvalidEncoding(goa.MissingPayloadError())
		}
		return ctrl.Push(rctx)
	}
	service.Mux.Handle("POST", "/api/push", ctrl.MuxHandler("Push", h, unmarshalPushOrdinalValuesPayload))
	service.LogInfo("mount", "ctrl", "OrdinalValues", "action", "Push", "route", "POST /api/push")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewRegisterOrdinalValuesContext(ctx, service)
		if err != nil {
			return err
		}
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*RegisterOrdinalValuesPayload)
		} else {
			return goa.ErrInvalidEncoding(goa.MissingPayloadError())
		}
		return ctrl.Register(rctx)
	}
	service.Mux.Handle("POST", "/api/register", ctrl.MuxHandler("Register", h, unmarshalRegisterOrdinalValuesPayload))
	service.LogInfo("mount", "ctrl", "OrdinalValues", "action", "Register", "route", "POST /api/register")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewStatisticsOrdinalValuesContext(ctx, service)
		if err != nil {
			return err
		}
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*StatisticsOrdinalValuesPayload)
		} else {
			return goa.ErrInvalidEncoding(goa.MissingPayloadError())
		}
		return ctrl.Statistics(rctx)
	}
	service.Mux.Handle("POST", "/api/statistics", ctrl.MuxHandler("Statistics", h, unmarshalStatisticsOrdinalValuesPayload))
	service.LogInfo("mount", "ctrl", "OrdinalValues", "action", "Statistics", "route", "POST /api/statistics")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewTagOrdinalValuesContext(ctx, service)
		if err != nil {
			return err
		}
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*TagOrdinalValuesPayload)
		} else {
			return goa.ErrInvalidEncoding(goa.MissingPayloadError())
		}
		return ctrl.Tag(rctx)
	}
	service.Mux.Handle("POST", "/api/tag", ctrl.MuxHandler("Tag", h, unmarshalTagOrdinalValuesPayload))
	service.LogInfo("mount", "ctrl", "OrdinalValues", "action", "Tag", "route", "POST /api/tag")
}

// unmarshalPushOrdinalValuesPayload unmarshals the request body into the context request data Payload field.
func unmarshalPushOrdinalValuesPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &pushOrdinalValuesPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalRegisterOrdinalValuesPayload unmarshals the request body into the context request data Payload field.
func unmarshalRegisterOrdinalValuesPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &registerOrdinalValuesPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalStatisticsOrdinalValuesPayload unmarshals the request body into the context request data Payload field.
func unmarshalStatisticsOrdinalValuesPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &statisticsOrdinalValuesPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalTagOrdinalValuesPayload unmarshals the request body into the context request data Payload field.
func unmarshalTagOrdinalValuesPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &tagOrdinalValuesPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}
