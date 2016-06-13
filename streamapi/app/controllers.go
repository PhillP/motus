//************************************************************************//
// API "Stream Statistics API": Application Controllers
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
}
