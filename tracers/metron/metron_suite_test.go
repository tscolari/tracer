package metron_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMetron(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Metron Suite")
}
