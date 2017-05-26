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
		tracer = new(metron.Tracer)
	})

	Describe("StartTransaction", func() {
		It("resets the tracer name", func() {
			Expect(tracer.Name()).To(Equal(""))
			tracer.StartTransaction("new!")
			Expect(tracer.Name()).To(Equal("new!"))
		})
	})

	Describe("StartSpan", func() {
		BeforeEach(func() {
			tracer.StartTransaction("test")
		})
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
