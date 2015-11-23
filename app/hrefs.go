//************************************************************************//
// goa Swagger service: Application Resource Href Factories
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/raphael/goa-swagger
// --design=github.com/raphael/goa-swagger/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "fmt"

// SpecHref returns the resource href.
func SpecHref(packagePath interface{}) string {
	return fmt.Sprintf("/swagger/spec/%v", packagePath)
}
