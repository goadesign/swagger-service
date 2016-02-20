//************************************************************************//
// API "goa Swagger service": Application Controllers
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/goadesign/swagger-service
// --design=github.com/goadesign/swagger-service/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"net/http"
)

// SpecController is the controller interface for the Spec actions.
type SpecController interface {
	goa.Controller
	Show(*ShowSpecContext) error
}

// MountSpecController "mounts" a Spec resource controller on the given service.
func MountSpecController(service goa.Service, ctrl SpecController) {
	// Setup encoders and decoders. This is idempotent and is done by each MountXXX function.
	service.SetEncoder(goa.GobEncoderFactory(), false, "application/gob", "application/x-gob")
	service.SetEncoder(goa.JSONEncoderFactory(), true, "application/json")
	service.SetEncoder(goa.XMLEncoderFactory(), false, "application/xml", "text/xml")
	service.SetDecoder(goa.GobDecoderFactory(), false, "application/gob", "application/x-gob")
	service.SetDecoder(goa.JSONDecoderFactory(), true, "application/json")
	service.SetDecoder(goa.XMLDecoderFactory(), false, "application/xml", "text/xml")

	// Setup endpoint handler
	var h goa.Handler
	mux := service.ServeMux()
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewShowSpecContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Show(rctx)
	}
	mux.Handle("GET", "/swagger/spec/*packagePath", ctrl.HandleFunc("Show", h, nil))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Spec"}, goa.KV{"action", "Show"}, goa.KV{"route", "GET /swagger/spec/*packagePath"})
}
