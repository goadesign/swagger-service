//************************************************************************//
// API "goa Swagger service": Application Controllers
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/goadesign/swagger-service/design
// --out=$(GOPATH)/src/github.com/goadesign/swagger-service
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"golang.org/x/net/context"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// AeController is the controller interface for the Ae actions.
type AeController interface {
	goa.Muxer
	Health(*HealthAeContext) error
	Start(*StartAeContext) error
}

// MountAeController "mounts" a Ae resource controller on the given service.
func MountAeController(service *goa.Service, ctrl AeController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewHealthAeContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Health(rctx)
	}
	service.Mux.Handle("GET", "/_ah/health", ctrl.MuxHandler("Health", h, nil))
	service.LogInfo("mount", "ctrl", "Ae", "action", "Health", "route", "GET /_ah/health")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewStartAeContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Start(rctx)
	}
	service.Mux.Handle("GET", "/_ah/start", ctrl.MuxHandler("Start", h, nil))
	service.LogInfo("mount", "ctrl", "Ae", "action", "Start", "route", "GET /_ah/start")
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
	service.Mux.Handle("OPTIONS", "/swagger/spec", ctrl.MuxHandler("preflight", handleSpecOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowSpecContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleSpecOrigin(h)
	service.Mux.Handle("GET", "/swagger/spec", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Spec", "action", "Show", "route", "GET /swagger/spec")
}

// handleSpecOrigin applies the CORS response headers corresponding to the origin.
func handleSpecOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}
