package main

import (
	"os"

	"github.com/raphael/goa"
	"github.com/raphael/goa-swagger/app"
	"github.com/raphael/goa-swagger/swagger"
	"github.com/raphael/goa/cors"
	"gopkg.in/inconshreveable/log15.v2"
)

func main() {
	// Configure logger
	goa.Log.SetHandler(log15.StreamHandler(os.Stderr, log15.LogfmtFormat()))

	// Create service
	service := goa.NewGraceful("goa Swagger service")

	// Setup CORS
	spec, _ := cors.New(func() {
		cors.Origin("*", func() {
			cors.Resource("*", func() {
				cors.Methods("GET")
			})
		})
	})

	// Setup middleware
	service.Use(goa.RequestID())
	service.Use(goa.LogRequest())
	service.Use(cors.Middleware(spec))
	service.Use(goa.Recover())

	// Mount "spec" controller
	c := NewSpecController(service)
	app.MountSpecController(service, c)

	// Mount Swagger spec provider controller
	swagger.MountController(service)

	// Start service, listen on port 8080
	if err := service.ListenAndServe(":8080"); err != nil {
		service.Crit(err.Error())
	}
}
