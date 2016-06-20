package stream

// Router sends data to an appropriate channel based on key 
type Router struct {
    channelMap              map[string]chan OrdinalValue
}

// NewRouter creates a router for an input channel and begins reading immediately
func NewRouter() (*Router) {
    var channelMap = make(map[string]chan OrdinalValue)
    
    router := Router {
        channelMap: channelMap}
    
    return &router
}

// Route values from an input channel
func (router *Router) Route(input chan OrdinalValue, unassigned chan OrdinalValue) {
    for v := range input {
        // lookup the appropriate channel based on key
        channel := router.channelMap[v.StreamKey]
    
        if channel != nil {
            channel <- v
        } else {
            // if a channel was not found add value to the unassigned channel
            unassigned <- v
        }
    }
    
    for _,v := range router.channelMap {
        close(v)
    }
    
}

// Register the channel for a stream based on key
func (router *Router) Register(streamKey string, c chan OrdinalValue) {
    router.channelMap[streamKey] = c
}
