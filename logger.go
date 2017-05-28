package tracer

import (
	"code.cloudfoundry.org/lager"
)

//go:generate counterfeiter code.cloudfoundry.org/lager.Logger

func NewLogger(logger lager.Logger, tracer Tracer) *Logger {
	return &Logger{
		logger: logger,
		tracer: tracer,
	}
}

type Logger struct {
	logger lager.Logger
	tracer Tracer
}

func (l *Logger) RegisterSink(sink lager.Sink) {
	l.logger.RegisterSink(sink)
}

func (l *Logger) Session(task string, data ...lager.Data) lager.Logger {
	tracer := l.tracer.StartSpan(task)
	logger := l.logger.Session(task, data...)

	return NewLogger(logger, tracer)
}

func (l *Logger) SessionName() string {
	return l.logger.SessionName()
}

func (l *Logger) Debug(action string, data ...lager.Data) {
	if isClosing(action) {
		defer l.tracer.End()
	}

	l.logger.Debug(action, data...)
}

func (l *Logger) Info(action string, data ...lager.Data) {
	if isClosing(action) {
		defer l.tracer.End()
	}

	l.logger.Info(action, data...)
}

func (l *Logger) Error(action string, err error, data ...lager.Data) {
	l.logger.Error(action, err, data...)
}

func (l *Logger) Fatal(action string, err error, data ...lager.Data) {
	l.logger.Fatal(action, err, data...)
}

func (l *Logger) WithData(data lager.Data) lager.Logger {
	l.logger = l.logger.WithData(data)
	return l
}
