package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
	"time"
)

// AddOrdinalValuesPath computes a request path to the add action of OrdinalValues.
func AddOrdinalValuesPath(stream string, ordinal int, value float64) string {
	return fmt.Sprintf("/api/add/%v/%v/%v", stream, ordinal, value)
}

// add a value to a stream referencing an ordinal position
func (c *Client) AddOrdinalValues(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewAddOrdinalValuesRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAddOrdinalValuesRequest create the request corresponding to the add action endpoint of the OrdinalValues resource.
func (c *Client) NewAddOrdinalValuesRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PushOrdinalValuesPayload is the OrdinalValues push action payload.
type PushOrdinalValuesPayload struct {
	// The ordinal position within the stream
	Ordinal int `json:"ordinal" xml:"ordinal"`
	// Identifies the stream that the ordinal value relates to
	Stream string `json:"stream" xml:"stream"`
	// The value at the ordinal position
	Value float64 `json:"value" xml:"value"`
}

// PushOrdinalValuesPath computes a request path to the push action of OrdinalValues.
func PushOrdinalValuesPath() string {
	return fmt.Sprintf("/api/push")
}

// Pushes a new ordinal value onto the stream
func (c *Client) PushOrdinalValues(ctx context.Context, path string, payload *PushOrdinalValuesPayload) (*http.Response, error) {
	req, err := c.NewPushOrdinalValuesRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewPushOrdinalValuesRequest create the request corresponding to the push action endpoint of the OrdinalValues resource.
func (c *Client) NewPushOrdinalValuesRequest(ctx context.Context, path string, payload *PushOrdinalValuesPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*") // Use default encoder
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// RegisterOrdinalValuesPayload is the OrdinalValues register action payload.
type RegisterOrdinalValuesPayload struct {
	// The ordinal position within the stream
	IntervalSize int `json:"intervalSize" xml:"intervalSize"`
	// The value at the ordinal position
	MaxIntervalLag int `json:"maxIntervalLag" xml:"maxIntervalLag"`
	// Identifies the stream that the definition relates to
	Stream string `json:"stream" xml:"stream"`
	// A set of tag values to be assigned to the stream
	Tags []string `json:"tags,omitempty" xml:"tags,omitempty"`
	// The value at the ordinal position
	TargetSampleSize int `json:"targetSampleSize" xml:"targetSampleSize"`
}

// RegisterOrdinalValuesPath computes a request path to the register action of OrdinalValues.
func RegisterOrdinalValuesPath() string {
	return fmt.Sprintf("/api/register")
}

// Registers a new stream
func (c *Client) RegisterOrdinalValues(ctx context.Context, path string, payload *RegisterOrdinalValuesPayload) (*http.Response, error) {
	req, err := c.NewRegisterOrdinalValuesRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewRegisterOrdinalValuesRequest create the request corresponding to the register action endpoint of the OrdinalValues resource.
func (c *Client) NewRegisterOrdinalValuesRequest(ctx context.Context, path string, payload *RegisterOrdinalValuesPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*") // Use default encoder
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// StatisticsOrdinalValuesPayload is the OrdinalValues statistics action payload.
type StatisticsOrdinalValuesPayload struct {
	// Specifies a maximum date time used to restrict the interval statistics returned.  Only statistics for intervals that are for a time range up until this date time value will be returned.
	MaxDateTime *time.Time `json:"maxDateTime,omitempty" xml:"maxDateTime,omitempty"`
	// Specifies a maximum ordinal value used to restrict the interval statistics returned.  Only statistics for intervals that end on or before this ordinal value will be returned.
	MaxOrdinal *int `json:"maxOrdinal,omitempty" xml:"maxOrdinal,omitempty"`
	// If true, results across multiple intervals will be merged together to produce a summary result.
	MergeIntervals *bool `json:"mergeIntervals,omitempty" xml:"mergeIntervals,omitempty"`
	// If true, results from multiple streams will be merged together to produce a summary result.
	MergeStreams *bool `json:"mergeStreams,omitempty" xml:"mergeStreams,omitempty"`
	// Specifies a minimum date time used to restrict the interval statistics returned.  Only statistics for intervals that are for a time range on or after this date time value will be returned.
	MinDateTime *time.Time `json:"minDateTime,omitempty" xml:"minDateTime,omitempty"`
	// Specifies a minimum ordinal value used to restrict the interval statistics returned.  Only statistics for intervals that begin on or after this ordinal value will be returned.
	MinOrdinal *int `json:"minOrdinal,omitempty" xml:"minOrdinal,omitempty"`
	// Specifies the criteria by which streams are to be matched
	StreamMatchCriteria *StreamMatchCriteria `json:"streamMatchCriteria,omitempty" xml:"streamMatchCriteria,omitempty"`
}

// StatisticsOrdinalValuesPath computes a request path to the statistics action of OrdinalValues.
func StatisticsOrdinalValuesPath() string {
	return fmt.Sprintf("/api/statistics")
}

// Gets statistics matching search criteria
func (c *Client) StatisticsOrdinalValues(ctx context.Context, path string, payload *StatisticsOrdinalValuesPayload) (*http.Response, error) {
	req, err := c.NewStatisticsOrdinalValuesRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewStatisticsOrdinalValuesRequest create the request corresponding to the statistics action endpoint of the OrdinalValues resource.
func (c *Client) NewStatisticsOrdinalValuesRequest(ctx context.Context, path string, payload *StatisticsOrdinalValuesPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*") // Use default encoder
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// TagOrdinalValuesPayload is the OrdinalValues tag action payload.
type TagOrdinalValuesPayload struct {
	// If true, previously assigned tags will be cleared
	ClearAll *bool `json:"clearAll,omitempty" xml:"clearAll,omitempty"`
	// Identifies the stream that the definition relates to
	Stream string `json:"stream" xml:"stream"`
	// An array of tags to be assigned
	TagsToAssign []string `json:"tagsToAssign,omitempty" xml:"tagsToAssign,omitempty"`
	// An array of tags to be unassigned
	TagsToUnassign []string `json:"tagsToUnassign,omitempty" xml:"tagsToUnassign,omitempty"`
}

// TagOrdinalValuesPath computes a request path to the tag action of OrdinalValues.
func TagOrdinalValuesPath() string {
	return fmt.Sprintf("/api/tag")
}

// Changes the tag assignments for a stream
func (c *Client) TagOrdinalValues(ctx context.Context, path string, payload *TagOrdinalValuesPayload) (*http.Response, error) {
	req, err := c.NewTagOrdinalValuesRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewTagOrdinalValuesRequest create the request corresponding to the tag action endpoint of the OrdinalValues resource.
func (c *Client) NewTagOrdinalValuesRequest(ctx context.Context, path string, payload *TagOrdinalValuesPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*") // Use default encoder
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	return req, nil
}
