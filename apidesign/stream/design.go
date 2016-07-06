package design

import (
        . "github.com/goadesign/goa/design"
        . "github.com/goadesign/goa/design/apidsl"
)

var _ = API("Stream Statistics API", func() {
        Title("Stream Statistics")
        Description("An API for stream statistics")
        Host("localhost:8080")
        Scheme("http")
        BasePath("api")
        Consumes("application/json")
        Produces("application/json")
})

//OrdinalValue encapsulates a value at an ordinal position for a stream
var OrdinalValue = Type("OrdinalValue", func() {
        Attribute("stream", String, "Identifies the stream that the ordinal value relates to")
        Attribute("ordinal", Integer, "The ordinal position within the stream")
        Attribute("value", Number, "The value at the ordinal position")
        Required("stream","ordinal","value")
})

//StreamDefinition specifies the interval configuration for a stream
var StreamDefinition = Type("StreamDefinition", func() {
        Attribute("stream", String, "Identifies the stream that the definition relates to")
        Attribute("intervalSize", Integer, "The ordinal position within the stream")
        Attribute("maxIntervalLag", Integer, "The value at the ordinal position")
        Attribute("targetSampleSize", Integer, "The value at the ordinal position")
        Attribute("tags", ArrayOf(String), "A set of tag values to be assigned to the stream")
        Required("stream","intervalSize","maxIntervalLag","targetSampleSize")
})

//StatisticsSearchCriteria specifies a set of criteria used to determine the set of statistics to be returned
var StatisticsSearchCriteria = Type("StatisticsSearchCriteria", func() {
        Attribute("streamMatchCriteria", StreamMatchCriteria, "Specifies the criteria by which streams are to be matched")
        Attribute("minOrdinal", Integer, "Specifies a minimum ordinal value used to restrict the interval statistics returned.  Only statistics for intervals that begin on or after this ordinal value will be returned.")
        Attribute("maxOrdinal", Integer, "Specifies a maximum ordinal value used to restrict the interval statistics returned.  Only statistics for intervals that end on or before this ordinal value will be returned.")
        Attribute("minDateTime", DateTime, "Specifies a minimum date time used to restrict the interval statistics returned.  Only statistics for intervals that are for a time range on or after this date time value will be returned.")
        Attribute("maxDateTime", DateTime, "Specifies a maximum date time used to restrict the interval statistics returned.  Only statistics for intervals that are for a time range up until this date time value will be returned.")
        Attribute("mergeStreams", Boolean, "If true, results from multiple streams will be merged together to produce a summary result.")
        Attribute("mergeIntervals", Boolean, "If true, results across multiple intervals will be merged together to produce a summary result.")
})

//StreamMatchCriteria specifies a set of criteria used to match streams
var StreamMatchCriteria = Type("StreamMatchCriteria", func() {
        Attribute("streamKeys", ArrayOf(String), "An optional array of streamKeys used to select streams")
        Attribute("includeWithAllTags", ArrayOf(String), "An optional array of tags. Streams tagged with all of these tags will be included")
        Attribute("excludeWithAnyTags", ArrayOf(String), "An optional array of tags.  Streams tagged with any of these tags will be excluded")
})

//TagAssignmentDefinition specifies a change in tag assignments for a stream
var TagAssignmentDefinition = Type("TagAssignmentDefinition", func() {
        Attribute("stream", String, "Identifies the stream that the definition relates to")
        Attribute("clearAll", Boolean, "If true, previously assigned tags will be cleared")
        Attribute("tagsToAssign", ArrayOf(String), "An array of tags to be assigned")
        Attribute("tagsToUnassign", ArrayOf(String), "An array of tags to be unassigned")
        Required("stream")
})

//IntervalStatisticsResult encapsulates the set of statistics calculated for an interval of a stream
var IntervalStatisticsResult = MediaType("vnd.application/goa.intervalstatisticsresult", func() {
    Description("A set of statistics based on the values of a stream for an interval")
    Attributes(func() {                              // Defines the media type attributes
        Attribute("streamKey", String, "identifies the stream for which the interval statistics have been derived")
        Attribute("intervalStart", Integer, "the ordinal position at the start of the interval")
        Attribute("intervalEnd", Integer, "the ordinal position at the end of the interval")
        Attribute("minimum", Number, "the minimum value occuring within the interval")
        Attribute("maximum", Number, "the maximum value occuring within the interval")
        Attribute("mean", Number, "the mean of the interval values")
        Attribute("count", Number, "the count of values occuring within the interval")
        Attribute("sum", Number, "the sum of values occuring within the interval")
        Attribute("sampleMean", Number, "the mean of the sample values")
        Attribute("sampleSum", Number, "the sum of sample values")
        Attribute("sampleCount", Number, "the count of sample values")
        Attribute("sampleStandardDeviation", Number, "the standard deviation of the values within the sample set")
        Attribute("coefficientOfVariation", Number, "a measure of the variability of values within the sample set")
    })
    Links(func() {             // Defines the links embedded in the media type
    })
    View("default", func() {   // Defines the default view
        Attribute("streamKey")     
        Attribute("intervalStart")     
        Attribute("intervalEnd")     
        Attribute("minimum")     
        Attribute("maximum")     
        Attribute("mean")     
        Attribute("count")     
        Attribute("sum")     
        Attribute("sampleMean")     
        Attribute("sampleSum")     
        Attribute("sampleCount")     
        Attribute("sampleStandardDeviation")     
        Attribute("coefficientOfVariation")     
    })
})

//StatisticsResults represents the results of a statistics query
var StatisticsResults = MediaType("vnd.application/goa.statisticsresults", func() {
    Description("The results of a statistics query")
    Attributes(func() {
        Attribute("intervalStatisticsList", ArrayOf(IntervalStatisticsResult), "A list of matching interval statistics") // Operation results attribute
    })
    Links(func() {
    })
    View("default", func() {
        Attribute("intervalStatisticsList")
    })
})

var _ = Resource("OrdinalValues", func() {
        Action("add", func() {
                Routing(GET("add/:stream/:ordinal/:value"))
                Description("add a value to a stream referencing an ordinal position")
                Params(func() {
                        Param("stream", String, "The stream for which the value is to be added")
                        Param("ordinal", Integer, "The ordinal position of the value")
                        Param("value", Number, "The value to be added to the stream")
                })
                Response(OK, "text/plain")
        })
        
        Action("push", func() {
                Routing(POST("/push"))
                Description("Pushes a new ordinal value onto the stream")
                Payload(OrdinalValue)
                Response(OK, "text/plain")
        })
        
        Action("register", func() {
                Routing(POST("/register"))
                Description("Registers a new stream")
                Payload(StreamDefinition)
                Response(OK, "text/plain")
        })
        
        Action("tag", func() {
                Routing(POST("/tag"))
                Description("Changes the tag assignments for a stream")
                Payload(TagAssignmentDefinition)
                Response(OK, "text/plain")
        })
        
        Action("statistics", func() {
                Routing(POST("/statistics"))
                Description("Gets statistics matching search criteria")
                Payload(StatisticsSearchCriteria)
                Response(OK, StatisticsResults)
        })
})