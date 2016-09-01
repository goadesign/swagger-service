package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/swagger-service/app"
)

// AeController implements the ae resource.
type AeController struct {
	*goa.Controller
}

// NewAeController creates a ae controller.
func NewAeController(service *goa.Service) *AeController {
	return &AeController{Controller: service.NewController("AeController")}
}

// Health runs the health action.
func (c *AeController) Health(ctx *app.HealthAeContext) error {
	return ctx.OK([]byte("ok"))
}

// Start runs the start action.
func (c *AeController) Start(ctx *app.StartAeContext) error {
	return ctx.OK([]byte("started"))
}
