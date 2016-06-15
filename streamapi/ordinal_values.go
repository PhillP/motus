package main

import (
	"github.com/goadesign/goa"
	"github.com/phillp/motus/streamapi/app"
	"github.com/phillp/motus/stream"
)

var inChannel = make(chan stream.OrdinalValue, 1000)
var outChannel = make(chan stream.IntervalStatistics, 1000)
var unassignedChannel = make(chan stream.OrdinalValue, 1000)
var isRouting = false
var router = stream.NewRouter()
var irDoneChannels []chan bool
    
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

// Push runs the push action.
func (c *OrdinalValuesController) Push(ctx *app.PushOrdinalValuesContext) error {
	if !isRouting {
		go router.Route(inChannel, unassignedChannel)
		isRouting = true
	}
	var ordinalValue = stream.NewOrdinalValue(ctx.Payload.Stream,int64(ctx.Payload.Ordinal),ctx.Payload.Value)
	inChannel <- ordinalValue
	
	return nil
}

// Register runs the register action.
func (c *OrdinalValuesController) Register(ctx *app.RegisterOrdinalValuesContext) error {
	if !isRouting {
		go router.Route(inChannel, unassignedChannel)
		isRouting = true
	}
	
	var intervalRouter = stream.NewIntervalRouter(ctx.Payload.Stream, int64(ctx.Payload.IntervalSize), stream.OrdinalInterval, uint32(ctx.Payload.MaxIntervalLag), uint32(ctx.Payload.TargetSampleSize))
    var streamInput = make(chan stream.OrdinalValue, 1000)
    var irDoneChannel = make(chan bool)
    irDoneChannels = append(irDoneChannels, irDoneChannel)
        
    go intervalRouter.AccumulateFromChannel(streamInput, outChannel, irDoneChannel)
    router.Register(ctx.Payload.Stream, streamInput)
	
	return nil
}
