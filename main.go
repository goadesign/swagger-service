package main

import (
	"os"
	"time"

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

	// Create service
	service := goa.New("goa Swagger service")
	service.UseLogger(goalog15.New(logger))
	server := goa.NewGraceful(service, false, time.Duration(0))

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
	c := NewSpecController(service)
	app.MountSpecController(service, c)

	// Mount Swagger spec provider controller
	swagger.MountController(service)

	// Start service, listen on port 8080
	if err := server.ListenAndServe(":8080"); err != nil {
		goa.Error(service.Context, err.Error())
	}
}
