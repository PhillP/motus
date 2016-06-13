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
