//************************************************************************//
// goa Swagger service Swagger Spec
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/raphael/goa-swagger
// --design=github.com/raphael/goa-swagger/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package swagger

import (
	"github.com/julienschmidt/httprouter"
	"github.com/raphael/goa"
)

// MountController mounts the swagger spec controller under "/swagger.json".
func MountController(service goa.Service) {
	ctrl := service.NewController("Swagger")
	service.Info("mount", "ctrl", "Swagger", "action", "Show", "route", "GET /swagger.json")
	h := ctrl.NewHTTPRouterHandle("Show", getSwagger)
	service.HTTPHandler().(*httprouter.Router).Handle("GET", "/swagger.json", h)
}

// getSwagger is the httprouter handle that returns the Swagger spec.
// func getSwagger(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
func getSwagger(ctx *goa.Context) error {
	ctx.Header().Set("Content-Type", "application/swagger+json")
	ctx.Header().Set("Cache-Control", "public, max-age=3600")
	return ctx.Respond(200, []byte(spec))
}

// Generated spec
const spec = `{"swagger":"2.0","info":{"title":"goa Swagger specification generation service","description":"The goa Swagger service can render the swagger specification of a goa design package given its Go path","contact":{"name":"The goa team","email":"admin@goa.design","url":"http://goa.design"},"license":{"name":"MIT","url":"https://github.com/raphael/goa/blob/master/LICENSE"},"version":""},"host":"swagger.goa.design","basePath":"/swagger","schemes":["http"],"consumes":["application/json"],"produces":["application/json"],"paths":{"/spec/{packagePath}":{"get":{"description":"Retrieve Swagger specification for given goa service design package","operationId":"spec#show","consumes":["application/json"],"produces":["application/json"],"parameters":[{"name":"packagePath","in":"path","description":"Go package path to goa service design package","required":true,"type":"string","format":"uri"}],"responses":{"200":{"description":""},"422":{"description":""}},"schemes":["https"]}}},"responses":{"UnprocessableEntity":{"description":""}},"externalDocs":{"description":"GoDoc","url":"https://godoc.org/github.com/raphael/goa-swagger"}} `
