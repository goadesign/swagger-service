//************************************************************************//
// goa Swagger service: Application Controllers
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/raphael/goa-swagger
// --design=github.com/raphael/goa-swagger/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/raphael/goa"
)

// SpecController is the controller interface for the Spec actions.
type SpecController interface {
	goa.Controller
	Show(*ShowSpecContext) error
}

// MountSpecController "mounts" a Spec resource controller on the given service.
func MountSpecController(service goa.Service, ctrl SpecController) {
	router := service.HTTPHandler().(*httprouter.Router)
	var h goa.Handler
	h = func(c *goa.Context) error {
		ctx, err := NewShowSpecContext(c)
		if err != nil {
			return goa.NewBadRequestError(err)
		}
		return ctrl.Show(ctx)
	}
	router.Handle("GET", "/swagger/spec/*packagePath", ctrl.NewHTTPRouterHandle("Show", h))
	service.Info("mount", "ctrl", "Spec", "action", "Show", "route", "GET /swagger/spec/*packagePath")
}
