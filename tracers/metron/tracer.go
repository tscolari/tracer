package metron

import (
	"errors"
	"time"

	"github.com/cloudfoundry/dropsonde/metrics"
)

func New() *Tracer {
	return &Tracer{}
}

type Tracer struct {
	start *time.Time
	name  string
}

func newTracer(id string) *Tracer {
	start := time.Now()

	return &Tracer{
		name:  id,
		start: &start,
	}
}

func (t *Tracer) StartTransaction(id string) {
	t.name = id
}

func (t *Tracer) StartSpan(name string) *Tracer {
	return newTracer(t.name + "." + name)
}

func (t *Tracer) End() error {
	if t.start == nil {
		return errors.New("tracer not initialized")
	}

	metrics.Value(t.name, float64(time.Since(*t.start).Nanoseconds()), "nanoseconds")
	return nil
}

func (t *Tracer) Name() string {
	return t.name
}

func (t *Tracer) StartedAt() time.Time {
	return *t.start
}
