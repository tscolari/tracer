package opentracing_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	ot "github.com/opentracing/opentracing-go"
	"github.com/tscolari/tracer/tracers/opentracing"
	"github.com/tscolari/tracer/tracers/opentracing/opentracingfakes"
)

var _ = Describe("Tracer", func() {
	var (
		otTracerImplementation *opentracingfakes.FakeTracer
		otSpanImplementation   *opentracingfakes.FakeSpan

		tracer *opentracing.Tracer
	)

	BeforeEach(func() {
		otTracerImplementation = new(opentracingfakes.FakeTracer)
		otSpanImplementation = new(opentracingfakes.FakeSpan)
		otTracerImplementation.StartSpanReturns(otSpanImplementation)
	})

	Describe("New", func() {
		BeforeEach(func() {
			tracer = opentracing.New("test", otTracerImplementation)
		})

		Describe("StartSpan", func() {
			It("uses the global opentracing Tracer", func() {
				tracer.StartSpan("hello")

				Expect(otTracerImplementation.StartSpanCallCount()).To(Equal(2))
				name, _ := otTracerImplementation.StartSpanArgsForCall(1)
				Expect(name).To(Equal("hello"))
			})
		})

		Describe("End", func() {
			It("uses the global opentracing Tracer", func() {
				tracer.End()

				Expect(otSpanImplementation.FinishCallCount()).To(Equal(1))
			})
		})
	})

	Describe("NewFromGlobal", func() {
		BeforeEach(func() {
			ot.InitGlobalTracer(otTracerImplementation)
			tracer = opentracing.NewFromGlobal("test")
		})

		Describe("StartSpan", func() {
			It("uses the global opentracing Tracer", func() {
				tracer.StartSpan("hello")

				Expect(otTracerImplementation.StartSpanCallCount()).To(Equal(2))
				name, _ := otTracerImplementation.StartSpanArgsForCall(1)
				Expect(name).To(Equal("hello"))
			})
		})

		Describe("End", func() {
			It("uses the global opentracing Tracer", func() {
				tracer.End()

				Expect(otSpanImplementation.FinishCallCount()).To(Equal(1))
			})
		})
	})

})
