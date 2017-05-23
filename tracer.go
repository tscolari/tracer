package tracer

//go:generate counterfeiter . Tracer
type Tracer interface {
	StartTransaction(id string)
	StartSpan(id string) Tracer
	End() error
}
