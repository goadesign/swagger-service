//************************************************************************//
// goa Swagger service: Application Contexts
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

import "github.com/raphael/goa"

// ShowSpecContext provides the spec show action context.
type ShowSpecContext struct {
	*goa.Context
	PackagePath string
}

// NewShowSpecContext parses the incoming request URL and body, performs validations and creates the
// context used by the spec controller show action.
func NewShowSpecContext(c *goa.Context) (*ShowSpecContext, error) {
	var err error
	ctx := ShowSpecContext{Context: c}
	rawPackagePath, ok := c.Get("packagePath")
	if ok {
		ctx.PackagePath = rawPackagePath
		if ctx.PackagePath != "" {
			if err2 := goa.ValidateFormat(goa.FormatURI, ctx.PackagePath); err2 != nil {
				err = goa.InvalidFormatError(`packagePath`, ctx.PackagePath, goa.FormatURI, err2, err)
			}
		}
	}
	return &ctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowSpecContext) OK(resp []byte) error {
	return ctx.Respond(200, resp)
}

// UnprocessableEntity sends a HTTP response with status code 422.
func (ctx *ShowSpecContext) UnprocessableEntity(resp []byte) error {
	return ctx.Respond(422, resp)
}
