package main

import (
	"os"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/logging/log15"
	"github.com/goadesign/goa/middleware"
	"github.com/goadesign/swagger-service/app"
	"gopkg.in/inconshreveable/log15.v2"
)

func main() {
	// Configure logger
	logger := log15.New()
	logger.SetHandler(log15.StreamHandler(os.Stderr, log15.LogfmtFormat()))

	// Create service
	service := goa.New("goa Swagger service")
	service.WithLogger(goalog15.New(logger))

	// Setup middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "spec" controller
	c := NewSpecController(service)
	app.MountSpecController(service, c)

	// Mount "swagger" controller
	app.MountSwaggerController(service, service.NewController("swagger"))

	// Start service, listen on port 8080
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError(err.Error())
	}
}
