package metron_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tscolari/tracer/tracers/metron"
)

var _ = Describe("Tracer", func() {
	var tracer *metron.Tracer

	BeforeEach(func() {
		tracer = metron.New("test")
	})

	Describe("StartSpan", func() {
		It("creates a new tracer with the new name", func() {
			newTracer := tracer.StartSpan("hello-world")
			Expect(newTracer.Name()).To(Equal("test.hello-world"))
		})

		It("marks the start time", func() {
			now := time.Now()
			newTracer := tracer.StartSpan("hello-world")

			Expect(newTracer.StartedAt().Unix()).To(BeNumerically("~", now.Unix(), 5))

		})
	})

})
