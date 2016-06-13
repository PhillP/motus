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

})