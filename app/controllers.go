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

// inited is true if initService has been called
var inited = false

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	if inited {
		return
	}
	inited = true

	// Setup encoders and decoders
	service.Encoder(goa.NewJSONEncoder, "application/json")
	service.Encoder(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder(goa.NewXMLEncoder, "application/xml")
	service.Decoder(goa.NewJSONDecoder, "application/json")
	service.Decoder(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder(goa.NewJSONEncoder, "*/*")
	service.Decoder(goa.NewJSONDecoder, "*/*")
}

// SpecController is the controller interface for the Spec actions.
type SpecController interface {
	goa.Muxer
	Show(*ShowSpecContext) error
}

// MountSpecController "mounts" a Spec resource controller on the given service.
func MountSpecController(service *goa.Service, ctrl SpecController) {
	initService(service)
	var h goa.Handler
	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewShowSpecContext(ctx)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/swagger/spec/*packagePath", ctrl.MuxHandler("Show", h, nil))
	goa.Info(goa.RootContext, "mount", goa.KV{"ctrl", "Spec"}, goa.KV{"action", "Show"}, goa.KV{"route", "GET /swagger/spec/*packagePath"})
}
