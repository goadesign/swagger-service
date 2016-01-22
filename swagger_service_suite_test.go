package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSwaggerService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SwaggerService Suite")
}
