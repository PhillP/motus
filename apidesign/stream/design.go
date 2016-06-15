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
        Required("stream","intervalSize","maxIntervalLag","targetSampleSize")
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

})