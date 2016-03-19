//************************************************************************//
// API "goa Swagger service": Application Contexts
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/goadesign/swagger-service
// --design=github.com/goadesign/swagger-service/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
)

// ShowSpecContext provides the spec show action context.
type ShowSpecContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	PackagePath string
}

// NewShowSpecContext parses the incoming request URL and body, performs validations and creates the
// context used by the spec controller show action.
func NewShowSpecContext(ctx context.Context) (*ShowSpecContext, error) {
	var err error
	req := goa.Request(ctx)
	rctx := ShowSpecContext{Context: ctx, ResponseData: goa.Response(ctx), RequestData: req}
	rawPackagePath := req.Params.Get("packagePath")
	if rawPackagePath != "" {
		rctx.PackagePath = rawPackagePath
		if err2 := goa.ValidateFormat(goa.FormatURI, rctx.PackagePath); err2 != nil {
			err = goa.StackErrors(err, goa.InvalidFormatError(`packagePath`, rctx.PackagePath, goa.FormatURI, err2))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowSpecContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/swagger")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// UnprocessableEntity sends a HTTP response with status code 422.
func (ctx *ShowSpecContext) UnprocessableEntity(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(422)
	_, err := ctx.ResponseData.Write(resp)
	return err
}
