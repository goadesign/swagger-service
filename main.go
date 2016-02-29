package main

import (
	"os"

	"github.com/goadesign/goa"
	"github.com/goadesign/logging/log15"
	"github.com/goadesign/middleware"
	"github.com/goadesign/middleware/cors"
	"github.com/goadesign/swagger-service/app"
	"github.com/goadesign/swagger-service/swagger"
	"gopkg.in/inconshreveable/log15.v2"
)

func main() {
	// Configure logger
	logger := log15.New()
	logger.SetHandler(log15.StreamHandler(os.Stderr, log15.LogfmtFormat()))
	goa.Log = goalog15.New(logger)

	// Create service
	service := goa.NewGraceful("goa Swagger service", false)

	// Setup CORS
	spec, _ := cors.New(func() {
		cors.Origin("*", func() {
			cors.Resource("*", func() {
				cors.Methods("GET")
			})
		})
	})

	// Setup middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(cors.Middleware(spec))
	service.Use(middleware.Recover())

	// Mount "spec" controller
	c := NewSpecController(service.Service)
	app.MountSpecController(service.Service, c)

	// Mount Swagger spec provider controller
	swagger.MountController(service.Service)

	// Start service, listen on port 8080
	if err := service.ListenAndServe(":8080"); err != nil {
		goa.Error(goa.RootContext, err.Error())
	}
}
