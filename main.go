package main

import (
	"github.com/raphael/goa"
	"github.com/raphael/goa-swagger/app"
	"github.com/raphael/goa-swagger/swagger"
	"github.com/raphael/goa/cors"
)

func main() {
	// Create service
	api := goa.NewGraceful("goa Swagger service")

	// Setup CORS
	spec, _ := cors.New(func() {
		cors.Origin("*", func() {
			cors.Resource("*", func() {
				cors.Methods("GET")
			})
		})
	})

	// Setup middleware
	api.Use(goa.RequestID())
	api.Use(goa.LogRequest())
	api.Use(cors.Middleware(spec))
	api.Use(goa.Recover())

	// Mount "spec" controller
	c := NewSpecController()
	app.MountSpecController(api, c)

	// Mount Swagger spec provider controller
	swagger.MountController(api)

	// Start service, listen on port 8080
	api.ListenAndServe(":8080")
}
