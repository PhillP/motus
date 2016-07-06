package main

import (
	"github.com/goadesign/goa"
	"github.com/phillp/motus/streamapi/app"
	"github.com/phillp/motus/stream"
	log "gopkg.in/inconshreveable/log15.v2"
)

var unassignedLog = log.New()
var resultsLog = log.New()
var inChannel = make(chan stream.OrdinalValue, 1000)
var outChannel = make(chan stream.IntervalStatistics, 1000)
var unassignedChannel = make(chan stream.OrdinalValue, 1000)
var isRouting = false
var router = stream.NewRouter()
var irDoneChannels []chan bool

func init() {
	var unassignedHandler, _ = log.FileHandler("./unassigned.log", log.JsonFormat())
	unassignedLog.SetHandler(unassignedHandler)
	
	var resultHandler, _ = log.FileHandler("./results.log", log.JsonFormat())
	resultsLog.SetHandler(resultHandler)
}
    
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
		// beginning logging unassigned values
		go logOrdinalValues("Stream not registered.", unassignedLog, unassignedChannel)
		
		// begin logging statistical results
		go logResults("Results for Interval.", resultsLog, outChannel)
		
		go router.Route(inChannel, unassignedChannel)
		isRouting = true
	}
	var ov = stream.NewOrdinalValue(ctx.Payload.Stream,int64(ctx.Payload.Ordinal),ctx.Payload.Value)
	inChannel <- ov
	
	return nil
}

func logOrdinalValues(message string, logger log.Logger, c chan stream.OrdinalValue) {
	for v := range c {
		logger.Info(message, "Stream", v.StreamKey, "Ordinal", v.Ordinal, "Value", v.Value)
	}
}

func logResults(message string, logger log.Logger, c chan stream.IntervalStatistics) {
	for v := range c {
		logger.Info(message, "Interval Start", v.IntervalStart, 
							 "Interval End", v.IntervalEnd,
							 "Count", v.Count,
							 "Minimum", v.Minimum,
							 "Maximum", v.Maximum,
							 "Mean", v.Mean,
							 "Sum", v.Sum,
							 "SampleCount", v.SampleCount,
							 "SampleMean", v.SampleMean,
							 "SampleStandardDeviation", v.SampleStandardDeviation,
							 "SampleSum", v.SampleSum,
							 "CoefficientOfVariation", v.CoefficientOfVariation)
		
	}
}

// Register runs the register action.
func (c *OrdinalValuesController) Register(ctx *app.RegisterOrdinalValuesContext) error {
	
	var intervalRouter = stream.NewIntervalRouter(ctx.Payload.Stream, int64(ctx.Payload.IntervalSize), stream.OrdinalInterval, uint32(ctx.Payload.MaxIntervalLag), uint32(ctx.Payload.TargetSampleSize))
    var streamInput = make(chan stream.OrdinalValue, 1000)
    var irDoneChannel = make(chan bool)
    irDoneChannels = append(irDoneChannels, irDoneChannel)
        
    go intervalRouter.AccumulateFromChannel(streamInput, outChannel, irDoneChannel)
    router.Register(ctx.Payload.Stream, streamInput)

	
	return nil
}

// Tag performs the action of applying tags to a stream
func (c *OrdinalValuesController) Tag(ctx *app.TagOrdinalValuesContext) error {
	return nil
}

// Statistics performs the action of calculating and returning statistics matching search criteria
func (c *OrdinalValuesController) Statistics(ctx *app.StatisticsOrdinalValuesContext) error {
	return nil
}


