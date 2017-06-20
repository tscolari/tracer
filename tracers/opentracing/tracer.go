package opentracing

import (
	"context"

	ot "github.com/opentracing/opentracing-go"
	"github.com/tscolari/tracer"
)

//go:generate counterfeiter github.com/opentracing/opentracing-go.Span
//go:generate counterfeiter github.com/opentracing/opentracing-go.Tracer

type Tracer struct {
	tracer ot.Tracer
	span   ot.Span
}

func New(name string, tracer ot.Tracer) *Tracer {
	return &Tracer{
		tracer: tracer,
		span:   tracer.StartSpan(name),
	}
}

func NewFromGlobal(name string) *Tracer {
	tracer := ot.GlobalTracer()
	return New(name, tracer)
}

func NewFromContext(ctx context.Context, name string) *Tracer {
	tracer := ot.GlobalTracer()
	span, _ := ot.StartSpanFromContext(ctx, name)
	return &Tracer{
		span:   span,
		tracer: tracer,
	}
}

func (t *Tracer) StartSpan(name string) tracer.Tracer {
	span := t.tracer.StartSpan(name, ot.ChildOf(t.span.Context()))

	return &Tracer{
		tracer: t.tracer,
		span:   span,
	}
}

func (t *Tracer) End() error {
	t.span.Finish()
	return nil
}
