package writer

import (
	"fmt"
	"io"
	"time"

	"github.com/pkg/errors"
)

func New(writer io.Writer) *Tracer {
	return &Tracer{
		writer: writer,
	}
}

type Tracer struct {
	start  *time.Time
	name   string
	writer io.Writer
}

func newTracer(id string, writer io.Writer) *Tracer {
	start := time.Now()

	return &Tracer{
		name:   id,
		start:  &start,
		writer: writer,
	}
}

func (t *Tracer) StartTransaction(id string) {
	t.name = id
}

func (t *Tracer) StartSpan(name string) *Tracer {
	return newTracer(t.name+"."+name, t.writer)
}

func (t *Tracer) End() error {
	if t.start == nil {
		return errors.New("tracer not initialized")
	}

	if _, err := t.writer.Write([]byte(traceLine(t.name, *t.start))); err != nil {
		return errors.Wrap(err, "writing trace line")
	}

	return nil
}

func (t *Tracer) Name() string {
	return t.name
}

func (t *Tracer) StartedAt() time.Time {
	return *t.start
}

func traceLine(name string, started time.Time) string {
	return fmt.Sprintf("%s: %d\n", name, time.Since(started).Nanoseconds())
}
