//************************************************************************//
// goa Swagger service Swagger Spec
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/goadesign/swagger-service
// --design=github.com/goadesign/swagger-service/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package swagger

import "github.com/goadesign/goa"

// MountController mounts the swagger spec controllers (one per API version) under "/swagger.json".
func MountController(service *goa.Service) {
	service.ServeFiles("/swagger.json", "swagger/swagger.json")

}
