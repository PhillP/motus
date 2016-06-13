package client

import (
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
