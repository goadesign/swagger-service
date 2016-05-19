package test

import (
	"bytes"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"github.com/goadesign/swagger-service/app"
	"golang.org/x/net/context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// ShowSpecOK test setup
func ShowSpecOK(t *testing.T, ctrl app.SpecController, packagePath string) {
	ShowSpecOKCtx(t, context.Background(), ctrl, packagePath)
}

// ShowSpecOKCtx test setup
func ShowSpecOKCtx(t *testing.T, ctx context.Context, ctrl app.SpecController, packagePath string) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/swagger/spec/*packagePath", packagePath), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["packagePath"] = []string{fmt.Sprintf("%v", packagePath)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "SpecTest"), rw, req, prms)
	showCtx, err := app.NewShowSpecContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

}

// ShowSpecUnprocessableEntity test setup
func ShowSpecUnprocessableEntity(t *testing.T, ctrl app.SpecController, packagePath string) {
	ShowSpecUnprocessableEntityCtx(t, context.Background(), ctrl, packagePath)
}

// ShowSpecUnprocessableEntityCtx test setup
func ShowSpecUnprocessableEntityCtx(t *testing.T, ctx context.Context, ctrl app.SpecController, packagePath string) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/swagger/spec/*packagePath", packagePath), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["packagePath"] = []string{fmt.Sprintf("%v", packagePath)}

	goaCtx := goa.NewContext(goa.WithAction(ctx, "SpecTest"), rw, req, prms)
	showCtx, err := app.NewShowSpecContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 422 {
		t.Errorf("invalid response status code: got %+v, expected 422", rw.Code)
	}

}
