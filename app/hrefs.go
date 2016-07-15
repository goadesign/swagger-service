//************************************************************************//
// API "goa Swagger service": Application Resource Href Factories
//
// Generated with goagen v0.2.dev, command line:
// $ goagen
// --design=github.com/goadesign/swagger-service/design
// --out=$(GOPATH)/src/github.com/goadesign/swagger-service
// --version=v0.2.dev
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "fmt"

// SpecHref returns the resource href.
func SpecHref(packagePath interface{}) string {
	return fmt.Sprintf("/swagger/spec/%v", packagePath)
}
