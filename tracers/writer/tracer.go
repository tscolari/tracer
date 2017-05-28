package writer

import (
	"fmt"
	"io"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func NewRand(writer io.Writer) *Tracer {
	name := uuid.New().String()
	return New(name, writer)
}

func New(name string, writer io.Writer) *Tracer {
	return &Tracer{
		name:   name,
		writer: writer,
	}
}

type Tracer struct {
	start  *time.Time
	name   string
	writer io.Writer
}

func (t *Tracer) StartSpan(name string) *Tracer {
	start := time.Now()

	return &Tracer{
		name:   t.name + "." + name,
		start:  &start,
		writer: t.writer,
	}
}

func (t *Tracer) End() error {
	if t.start == nil {
		return errors.New("tracer not initialized")
	}

	if _, err := t.writer.Write([]byte(t.traceLine())); err != nil {
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

func (t *Tracer) traceLine() string {
	return fmt.Sprintf("%s: %d\n", t.name, time.Since(*t.start).Nanoseconds())
}
