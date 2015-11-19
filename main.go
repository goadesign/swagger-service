package main

import (
	"github.com/raphael/goa"
	"github.com/raphael/goa-swagger/app"
	"github.com/raphael/goa-swagger/swagger"
)

func main() {
	// Create service
	api := goa.New("goa Swagger service")

	// Setup middleware
	api.Use(goa.RequestID())
	api.Use(goa.LogRequest())
	api.Use(goa.Recover())

	// Mount "spec" controller
	c := NewSpecController()
	app.MountSpecController(api, c)

	// Mount Swagger spec provider controller
	swagger.MountController(api)

	// Start service, listen on port 8080
	api.ListenAndServe(":8080")
}
