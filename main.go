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
	if err := api.ListenAndServe(":8080"); err != nil {
		api.Crit(err.Error())
	}
}
