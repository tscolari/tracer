package tracer_test

import (
	"errors"

	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/lager/lagertest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tscolari/tracer"
	"github.com/tscolari/tracer/tracerfakes"
)

var _ = Describe("Logger", func() {
	var (
		fakeLogger *lagertest.TestLogger
		fakeTracer *tracerfakes.FakeTracer
		logger     *tracer.Logger
	)

	BeforeEach(func() {
		fakeLogger = lagertest.NewTestLogger("tracer-test")
		fakeTracer = new(tracerfakes.FakeTracer)
		fakeTracer.StartSpanReturns(fakeTracer)

		logger = tracer.NewLogger(fakeLogger, fakeTracer)
	})

	Describe("it works as a normal logger", func() {
		It("logs DEBUG messages", func() {
			logger.Debug("hello", lager.Data{"key": "value"})

			Expect(fakeLogger.Logs()).To(HaveLen(1))
			logEntry := fakeLogger.Logs()[0]

			Expect(logEntry.Message).To(Equal("tracer-test.hello"))
			Expect(logEntry.Data).To(Equal(lager.Data{"key": "value"}))
			Expect(logEntry.LogLevel).To(Equal(lager.DEBUG))
		})

		It("logs INFO messages", func() {
			logger.Info("hello", lager.Data{"key": "value"})

			Expect(fakeLogger.Logs()).To(HaveLen(1))
			logEntry := fakeLogger.Logs()[0]

			Expect(logEntry.Message).To(Equal("tracer-test.hello"))
			Expect(logEntry.Data).To(Equal(lager.Data{"key": "value"}))
			Expect(logEntry.LogLevel).To(Equal(lager.INFO))
		})

		It("logs ERROR messages", func() {
			err := errors.New("failed")
			logger.Error("hello", err, lager.Data{"key": "value"})

			Expect(fakeLogger.Logs()).To(HaveLen(1))
			logEntry := fakeLogger.Logs()[0]

			Expect(logEntry.Message).To(Equal("tracer-test.hello"))
			Expect(logEntry.Data).To(Equal(lager.Data{"key": "value", "error": err.Error()}))
			Expect(logEntry.LogLevel).To(Equal(lager.ERROR))
		})
	})

	Describe("Session", func() {
		It("creates a new logger with the new session inside", func() {
			newLogger := logger.Session("new", lager.Data{"data": true})
			Expect(newLogger.SessionName()).To(Equal("tracer-test.new"))
			Expect(fakeTracer.StartSpanCallCount()).To(Equal(1))
		})

		It("creates a new logger with the new tracer span inside", func() {
			_ = logger.Session("new", lager.Data{"data": true})
			Expect(fakeTracer.StartSpanCallCount()).To(Equal(1))
			id := fakeTracer.StartSpanArgsForCall(0)
			Expect(id).To(Equal("new"))
		})
	})

	Describe("Trace End", func() {
		Describe("Debug", func() {
			It("ends the tracing when closing words are found", func() {
				logger.Debug("end")
				Expect(fakeTracer.EndCallCount()).To(Equal(1))
			})
		})

		Describe("Info", func() {
			It("ends the tracing when closing words are found", func() {
				logger.Info("end")
				Expect(fakeTracer.EndCallCount()).To(Equal(1))
			})
		})
	})
})
