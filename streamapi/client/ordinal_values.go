package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
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
