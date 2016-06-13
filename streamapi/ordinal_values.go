package main

import (
	"github.com/goadesign/goa"
	"github.com/phillp/motus/streamapi/app"
	"github.com/phillp/motus/stream"
)

var inChannel = make(chan stream.OrdinalValue)
var unassignedChannel = make(chan stream.OrdinalValue)
var isRouting = false
var router = stream.NewRouter()

// OrdinalValuesController implements the OrdinalValues resource.
type OrdinalValuesController struct {
	*goa.Controller
}

// NewOrdinalValuesController creates a OrdinalValues controller.
func NewOrdinalValuesController(service *goa.Service) *OrdinalValuesController {
	return &OrdinalValuesController{Controller: service.NewController("OrdinalValuesController")}
}

// Add runs the add action.
func (c *OrdinalValuesController) Add(ctx *app.AddOrdinalValuesContext) error {
	if !isRouting {
		go router.Route(inChannel, unassignedChannel)
		isRouting = true
	}
	var ordinalValue = stream.NewOrdinalValue(ctx.Stream,int64(ctx.Ordinal),ctx.Value)
	inChannel <- ordinalValue
	
	return nil
}
