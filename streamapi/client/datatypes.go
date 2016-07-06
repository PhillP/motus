//************************************************************************//
// User Types
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --design=github.com/phillp/motus/apidesign/stream
// --out=streamapi
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package client

import "net/http"

// StreamMatchCriteria user type.
type StreamMatchCriteria struct {
	// An optional array of tags.  Streams tagged with any of these tags will be excluded
	ExcludeWithAnyTags []string `json:"excludeWithAnyTags,omitempty" xml:"excludeWithAnyTags,omitempty"`
	// An optional array of tags. Streams tagged with all of these tags will be included
	IncludeWithAllTags []string `json:"includeWithAllTags,omitempty" xml:"includeWithAllTags,omitempty"`
	// An optional array of streamKeys used to select streams
	StreamKeys []string `json:"streamKeys,omitempty" xml:"streamKeys,omitempty"`
}

// The results of a statistics query
type GoaStatisticsresults struct {
	// A list of matching interval statistics
	IntervalStatisticsList []*GoaIntervalstatisticsresult `json:"intervalStatisticsList,omitempty" xml:"intervalStatisticsList,omitempty"`
}

// DecodeGoaStatisticsresults decodes the GoaStatisticsresults instance encoded in resp body.
func (c *Client) DecodeGoaStatisticsresults(resp *http.Response) (*GoaStatisticsresults, error) {
	var decoded GoaStatisticsresults
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// A set of statistics based on the values of a stream for an interval
type GoaIntervalstatisticsresult struct {
	// a measure of the variability of values within the sample set
	CoefficientOfVariation *float64 `json:"coefficientOfVariation,omitempty" xml:"coefficientOfVariation,omitempty"`
	// the count of values occuring within the interval
	Count *float64 `json:"count,omitempty" xml:"count,omitempty"`
	// the ordinal position at the end of the interval
	IntervalEnd *int `json:"intervalEnd,omitempty" xml:"intervalEnd,omitempty"`
	// the ordinal position at the start of the interval
	IntervalStart *int `json:"intervalStart,omitempty" xml:"intervalStart,omitempty"`
	// the maximum value occuring within the interval
	Maximum *float64 `json:"maximum,omitempty" xml:"maximum,omitempty"`
	// the mean of the interval values
	Mean *float64 `json:"mean,omitempty" xml:"mean,omitempty"`
	// the minimum value occuring within the interval
	Minimum *float64 `json:"minimum,omitempty" xml:"minimum,omitempty"`
	// the count of sample values
	SampleCount *float64 `json:"sampleCount,omitempty" xml:"sampleCount,omitempty"`
	// the mean of the sample values
	SampleMean *float64 `json:"sampleMean,omitempty" xml:"sampleMean,omitempty"`
	// the standard deviation of the values within the sample set
	SampleStandardDeviation *float64 `json:"sampleStandardDeviation,omitempty" xml:"sampleStandardDeviation,omitempty"`
	// the sum of sample values
	SampleSum *float64 `json:"sampleSum,omitempty" xml:"sampleSum,omitempty"`
	// identifies the stream for which the interval statistics have been derived
	StreamKey *string `json:"streamKey,omitempty" xml:"streamKey,omitempty"`
	// the sum of values occuring within the interval
	Sum *float64 `json:"sum,omitempty" xml:"sum,omitempty"`
}
