package test

import (
	"bytes"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"github.com/phillp/motus/streamapi/app"
	"golang.org/x/net/context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// AddOrdinalValuesOK test setup
func AddOrdinalValuesOK(t *testing.T, ctrl app.OrdinalValuesController, stream string, ordinal int, value float64) {
	AddOrdinalValuesOKCtx(t, context.Background(), ctrl, stream, ordinal, value)
}

// AddOrdinalValuesOKCtx test setup
func AddOrdinalValuesOKCtx(t *testing.T, ctx context.Context, ctrl app.OrdinalValuesController, stream string, ordinal int, value float64) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/add/%v/%v/%v", stream, ordinal, value), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["stream"] = []string{fmt.Sprintf("%v", stream)}
	prms["ordinal"] = []string{fmt.Sprintf("%v", ordinal)}
	prms["value"] = []string{fmt.Sprintf("%v", value)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "OrdinalValuesTest"), rw, req, prms)
	addCtx, err := app.NewAddOrdinalValuesContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}

	err = ctrl.Add(addCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

}

// PushOrdinalValuesOK test setup
func PushOrdinalValuesOK(t *testing.T, ctrl app.OrdinalValuesController, payload *app.PushOrdinalValuesPayload) {
	PushOrdinalValuesOKCtx(t, context.Background(), ctrl, payload)
}

// PushOrdinalValuesOKCtx test setup
func PushOrdinalValuesOKCtx(t *testing.T, ctx context.Context, ctrl app.OrdinalValuesController, payload *app.PushOrdinalValuesPayload) {
	err := payload.Validate()
	if err != nil {
		e, ok := err.(*goa.Error)
		if !ok {
			panic(err) //bug
		}
		if e.Status != 200 {
			t.Errorf("unexpected payload validation error: %+v", e)
		}
		return
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("POST", fmt.Sprintf("/api/push"), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "OrdinalValuesTest"), rw, req, prms)
	pushCtx, err := app.NewPushOrdinalValuesContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	pushCtx.Payload = payload

	err = ctrl.Push(pushCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

}

// RegisterOrdinalValuesOK test setup
func RegisterOrdinalValuesOK(t *testing.T, ctrl app.OrdinalValuesController, payload *app.RegisterOrdinalValuesPayload) {
	RegisterOrdinalValuesOKCtx(t, context.Background(), ctrl, payload)
}

// RegisterOrdinalValuesOKCtx test setup
func RegisterOrdinalValuesOKCtx(t *testing.T, ctx context.Context, ctrl app.OrdinalValuesController, payload *app.RegisterOrdinalValuesPayload) {
	err := payload.Validate()
	if err != nil {
		e, ok := err.(*goa.Error)
		if !ok {
			panic(err) //bug
		}
		if e.Status != 200 {
			t.Errorf("unexpected payload validation error: %+v", e)
		}
		return
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("POST", fmt.Sprintf("/api/register"), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "OrdinalValuesTest"), rw, req, prms)
	registerCtx, err := app.NewRegisterOrdinalValuesContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	registerCtx.Payload = payload

	err = ctrl.Register(registerCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

}

// StatisticsOrdinalValuesOK test setup
func StatisticsOrdinalValuesOK(t *testing.T, ctrl app.OrdinalValuesController, payload *app.StatisticsOrdinalValuesPayload) *app.GoaStatisticsresults {
	return StatisticsOrdinalValuesOKCtx(t, context.Background(), ctrl, payload)
}

// StatisticsOrdinalValuesOKCtx test setup
func StatisticsOrdinalValuesOKCtx(t *testing.T, ctx context.Context, ctrl app.OrdinalValuesController, payload *app.StatisticsOrdinalValuesPayload) *app.GoaStatisticsresults {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("POST", fmt.Sprintf("/api/statistics"), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "OrdinalValuesTest"), rw, req, prms)
	statisticsCtx, err := app.NewStatisticsOrdinalValuesContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	statisticsCtx.Payload = payload

	err = ctrl.Statistics(statisticsCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.GoaStatisticsresults)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.GoaStatisticsresults", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	return a

}

// TagOrdinalValuesOK test setup
func TagOrdinalValuesOK(t *testing.T, ctrl app.OrdinalValuesController, payload *app.TagOrdinalValuesPayload) {
	TagOrdinalValuesOKCtx(t, context.Background(), ctrl, payload)
}

// TagOrdinalValuesOKCtx test setup
func TagOrdinalValuesOKCtx(t *testing.T, ctx context.Context, ctrl app.OrdinalValuesController, payload *app.TagOrdinalValuesPayload) {
	err := payload.Validate()
	if err != nil {
		e, ok := err.(*goa.Error)
		if !ok {
			panic(err) //bug
		}
		if e.Status != 200 {
			t.Errorf("unexpected payload validation error: %+v", e)
		}
		return
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("POST", fmt.Sprintf("/api/tag"), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "OrdinalValuesTest"), rw, req, prms)
	tagCtx, err := app.NewTagOrdinalValuesContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	tagCtx.Payload = payload

	err = ctrl.Tag(tagCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

}
