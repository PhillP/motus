package stream

import (
    "math/rand"
    "testing"
    log "github.com/cihub/seelog"
    //"github.com/stretchr/testify/assert"
    "strconv"
    "fmt"
)

func TestRouter(t *testing.T) {
    var outchannel = make(chan IntervalStatistics)
    var inchannel = make(chan OrdinalValue)
    var unassignedchannel = make(chan OrdinalValue)
    var doneChannel = make(chan bool)
    
    // start logging the output statistics
    go ProcessOutputStatistics(outchannel, doneChannel)
    
    // make a router and start routing
    var router = NewRouter()
    go router.Route(inchannel, unassignedchannel)
    
    var irDoneChannels []chan bool
    
    for s:=1;s<=3;s++ {
        var streamName = "Stream" + strconv.Itoa(s)
        var randSource = rand.NewSource(int64(s))
        var r = rand.New(randSource)
        
        var intervalRouter = NewIntervalRouter(streamName, 10000, OrdinalInterval, 3, 500)
        var streamInput = make(chan OrdinalValue)
        var irDoneChannel = make(chan bool)
        irDoneChannels = append(irDoneChannels, irDoneChannel)
        
        go intervalRouter.AccumulateFromChannel(streamInput, outchannel, irDoneChannel)
        router.Register(streamName, streamInput)
        
        // start generating input values
        for i:=0;i<100000;i++ {
            var ov = NewOrdinalValue(streamName,int64(i),r.Float64())
            inchannel <- ov
        }
    }
    
    close(inchannel)
    close(unassignedchannel)
    
    for _,v := range irDoneChannels {
        <- v
        close(v)
    }
    
    close(outchannel)
    <- doneChannel
    close(doneChannel)
}

func ProcessOutputStatistics(output chan IntervalStatistics, done chan bool) {
    defer log.Flush()

    testConfig := `
    <seelog>
    <outputs>
    <rollingfile type="size" filename="./testlogs/roll.log" maxsize="1000000" maxrolls="1" />
    </outputs>
    </seelog>
    `
    logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))
    log.ReplaceLogger(logger)

    log.Info("-----------------------------Starting------------------------------")
    
    for s := range output {
        log.Info(fmt.Sprintf("Interval: %d to %d, minimum: %f, maximum: %f, mean: %f, sd: %f", s.intervalStart, s.intervalEnd, s.minimum, s.maximum, s.mean, s.sampleStandardDeviation))
    }
    
    log.Info("------------------------------Done--------------------------------")
    done <- true
}