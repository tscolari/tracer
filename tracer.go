package tracer

//go:generate counterfeiter . Tracer
type Tracer interface {
	StartSpan(id string) Tracer
	End() error
}
