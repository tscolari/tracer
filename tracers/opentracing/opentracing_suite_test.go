package opentracing_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestOpentracing(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Opentracing Suite")
}
