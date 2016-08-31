//************************************************************************//
// API "goa Swagger service": Application Resource Href Factories
//
// Generated with goagen v1.0.0, command line:
// $ goagen
// --design=github.com/goadesign/swagger-service/design
// --out=$(GOPATH)/src/github.com/goadesign/swagger-service
// --version=v1.0.0
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"fmt"
	"strings"
)

// SpecHref returns the resource href.
func SpecHref(packagePath interface{}) string {
	parampackagePath := strings.TrimLeftFunc(fmt.Sprintf("%v", packagePath), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/swagger/spec/%v", parampackagePath)
}
