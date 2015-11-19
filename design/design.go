package design

import (
	. "github.com/raphael/goa/design"
	. "github.com/raphael/goa/design/dsl"
)

var _ = API("goa Swagger service", func() {
	Title("goa Swagger specification generation service")
	Description("The goa Swagger service can render the swagger specification of a goa design package given its Go path")
	Contact(func() {
		Name("The goa team")
		Email("admin@goa.design")
		URL("http://goa.design")
	})
	License(func() {
		Name("MIT")
		URL("https://github.com/raphael/goa/blob/master/LICENSE")
	})
	Docs(func() {
		Description("GoDoc")
		URL("https://godoc.org/github.com/raphael/goa-swagger")
	})
	Host("swagger.goa.design")
	Schemes("http")
	BasePath("/swagger")
	ResponseTemplate(UnprocessableEntity, func() {
		Media("text/plain")
		Status(422)
	})
})

var _ = Resource("spec", func() {
	DefaultMedia("application/swagger+json")
	BasePath("/spec")

	Action("show", func() {
		Routing(GET("/*packagePath"))
		Description("Retrieve Swagger specification for given goa service design package")
		Params(func() {
			Param("packagePath", String, "Go package path to goa service design package", func() {
				Format("uri")
			})
		})
		Response(OK)
		Response(UnprocessableEntity)
	})
})
