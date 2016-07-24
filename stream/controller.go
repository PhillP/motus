package stream

import (
    "time"
   	log "gopkg.in/inconshreveable/log15.v2"
)

// Controller is used to dispatch all operations related to stream processing
type Controller struct {
    contextKey              string
    
    router                  *Router
    tagRepository           TagRepository
    statisticsCache         CachedStatisticsRepository
    inChannel               chan OrdinalValue
    unassignedChannel       chan OrdinalValue
    outChannel              chan IntervalStatistics
    irDoneChannels          []chan bool
    
    unassignedLog           log.Logger
    resultsLog              log.Logger
}

//NewController creates a controller used to manage streams
func NewController(contextKey string) *Controller {
    var ulog = log.New()
    var rlog = log.New()

    /*
	var unassignedHandler, _ = log.FileHandler("./unassigned.log", log.JsonFormat())
	unassignedLog.SetHandler(unassignedHandler)
	
	var resultHandler, _ = log.FileHandler("./results.log", log.JsonFormat())
	resultsLog.SetHandler(resultHandler)
    */
    
    controller := Controller {
        contextKey: contextKey,
        router: NewRouter(),
        tagRepository: NewTagRepository(),
        statisticsCache: NewCachedStatisticsRepository(),
        inChannel: make(chan OrdinalValue, 1000),
        unassignedChannel: make(chan OrdinalValue, 1000),
        outChannel: make(chan IntervalStatistics, 1000),
        irDoneChannels: make([]chan bool, 0),
        unassignedLog: ulog,
        resultsLog: rlog}
        
     return &controller
}
  
  
// Register a stream within the controller
func (controller *Controller) Register(params *RegisterParams) (err error) {
	var intervalRouter = NewIntervalRouter(params.Stream, int64(params.IntervalSize), OrdinalInterval, uint32(params.MaxIntervalLag), uint32(params.TargetSampleSize))
    var streamInput = make(chan OrdinalValue, 1000)
    var irDoneChannel = make(chan bool)
    controller.irDoneChannels = append(controller.irDoneChannels, irDoneChannel)
        
    go intervalRouter.AccumulateFromChannel(streamInput, controller.outChannel, irDoneChannel)
    controller.router.Register(params.Stream, streamInput)

	return nil
}

// Unregister streams from the controller
func (controller *Controller) Unregister(params *UnregisterParams) (err error) {
    go func () {
        controller.router.StopAndUnregister(params.Stream)
        controller.ClearStatisticsCache(&ClearParams{ Stream: params.Stream })
    }()
    
    return nil
}

// Push new values onto a stream
func (controller *Controller) Push(params *PushParams) (err error) {
    
    controller.inChannel <- OrdinalValue {
       StreamKey: params.Stream,
       Ordinal: int64(params.Ordinal),
       Value: params.Value}
    
    return nil
}

// GetStatistics returns statistics matching search parameters
func (controller *Controller) GetStatistics(params *StatisticsParams) (err error) {
 return nil
}

// ModifyTags modifies the tags associated with a stream
func (controller *Controller) ModifyTags(params *TagParams) (err error) {
    if len(params.TagsToAssign) > 0 || params.ClearAll {
        controller.tagRepository.ApplyTags(params.Stream, params.TagsToAssign, params.ClearAll)
    }
    
    if len(params.TagsToUnassign) > 0 {
        controller.tagRepository.RemoveTags(params.Stream, params.TagsToUnassign)
    }
    
    return nil
}

// ClearStatisticsCache removes existing statistics data from cache
func (controller *Controller) ClearStatisticsCache(params *ClearParams) (err error) {
return nil
}

// PushParams encapsulates the information required for the push action
type PushParams struct {
	// The ordinal position within the stream
	Ordinal int `json:"ordinal,omitempty" xml:"ordinal,omitempty"`
	// Identifies the stream that the ordinal value relates to
	Stream string `json:"stream,omitempty" xml:"stream,omitempty"`
	// The value at the ordinal position
	Value float64 `json:"value,omitempty" xml:"value,omitempty"`
}

// RegisterParams encapsulates the information required for the register action
type RegisterParams struct {
	// Identifies the stream that the definition relates to
	Stream string `json:"stream,omitempty" xml:"stream,omitempty"`
	// The ordinal position within the stream
	IntervalSize int `json:"intervalSize,omitempty" xml:"intervalSize,omitempty"`
	// The value at the ordinal position
	MaxIntervalLag int `json:"maxIntervalLag,omitempty" xml:"maxIntervalLag,omitempty"`
	// A set of tag values to be assigned to the stream
	Tags []string `json:"tags,omitempty" xml:"tags,omitempty"`
	// The value at the ordinal position
	TargetSampleSize int `json:"targetSampleSize,omitempty" xml:"targetSampleSize,omitempty"`
}

// streamMatchSearchParams defines parameters used to match streams
type streamMatchSearchParams struct {
	// An optional array of tags.  Streams tagged with any of these tags will be excluded
	ExcludeWithAnyTags []string `json:"excludeWithAnyTags,omitempty" xml:"excludeWithAnyTags,omitempty"`
	// An optional array of tags. Streams tagged with all of these tags will be included
	IncludeWithAllTags []string `json:"includeWithAllTags,omitempty" xml:"includeWithAllTags,omitempty"`
	// An optional array of streamKeys used to select streams
	StreamKeys []string `json:"streamKeys,omitempty" xml:"streamKeys,omitempty"`
}

// StatisticsParams encapsulates the information required for the GetStatistics action
type StatisticsParams struct {
	// Specifies a maximum date time used to restrict the interval statistics returned.  Only statistics for intervals that are for a time range up until this date time value will be returned.
	MaxDateTime time.Time `json:"maxDateTime,omitempty" xml:"maxDateTime,omitempty"`
	// Specifies a maximum ordinal value used to restrict the interval statistics returned.  Only statistics for intervals that end on or before this ordinal value will be returned.
	MaxOrdinal int `json:"maxOrdinal,omitempty" xml:"maxOrdinal,omitempty"`
	// If true, results across multiple intervals will be merged together to produce a summary result.
	MergeIntervals bool `json:"mergeIntervals,omitempty" xml:"mergeIntervals,omitempty"`
	// If true, results from multiple streams will be merged together to produce a summary result.
	MergeStreams bool `json:"mergeStreams,omitempty" xml:"mergeStreams,omitempty"`
	// Specifies a minimum date time used to restrict the interval statistics returned.  Only statistics for intervals that are for a time range on or after this date time value will be returned.
	MinDateTime time.Time `json:"minDateTime,omitempty" xml:"minDateTime,omitempty"`
	// Specifies a minimum ordinal value used to restrict the interval statistics returned.  Only statistics for intervals that begin on or after this ordinal value will be returned.
	MinOrdinal int `json:"minOrdinal,omitempty" xml:"minOrdinal,omitempty"`
	// Specifies the criteria by which streams are to be matched
	StreamMatchSearchParams streamMatchSearchParams `json:"streamMatchCriteria,omitempty" xml:"streamMatchCriteria,omitempty"`
}

// TagParams encapsulates the information required for the ModifyTags action
type TagParams struct {
	// If true, previously assigned tags will be cleared
	ClearAll bool `json:"clearAll,omitempty" xml:"clearAll,omitempty"`
	// Identifies the stream that the definition relates to
	Stream string `json:"stream,omitempty" xml:"stream,omitempty"`
	// An array of tags to be assigned
	TagsToAssign []string `json:"tagsToAssign,omitempty" xml:"tagsToAssign,omitempty"`
	// An array of tags to be unassigned
	TagsToUnassign []string `json:"tagsToUnassign,omitempty" xml:"tagsToUnassign,omitempty"`
}

// ClearParams encapsulates the information required for the clear action
type ClearParams struct {
	// Identifies the stream
	Stream string `json:"stream,omitempty" xml:"stream,omitempty"`
    // If true, all data for the stream(s) will be cleared
	ClearAll bool `json:"clearAll,omitempty" xml:"clearAll,omitempty"`
	// An array of tags to be matched
	WithTags []string `json:"tagsToAssign,omitempty" xml:"tagsToAssign,omitempty"`
	// An array of tags that exclude a stream from a match
	ExcludingTags []string `json:"tagsToUnassign,omitempty" xml:"tagsToUnassign,omitempty"`
	// Specifies a maximum ordinal value used to restrict the clear operation.
	MaxOrdinal int `json:"maxOrdinal,omitempty" xml:"maxOrdinal,omitempty"`
}

// UnregisterParams encapsulates the information required for the unregister action
type UnregisterParams struct {
	// Identifies the stream
	Stream string `json:"stream,omitempty" xml:"stream,omitempty"`
	// An array of tags to be matched
	WithTags []string `json:"tagsToAssign,omitempty" xml:"tagsToAssign,omitempty"`
	// An array of tags that exclude a stream from a match
	ExcludingTags []string `json:"tagsToUnassign,omitempty" xml:"tagsToUnassign,omitempty"`
}