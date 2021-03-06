// This file was generated by counterfeiter
package opentracingfakes

import (
	"sync"

	opentracing "github.com/opentracing/opentracing-go"
)

type FakeTracer struct {
	StartSpanStub        func(operationName string, opts ...opentracing.StartSpanOption) opentracing.Span
	startSpanMutex       sync.RWMutex
	startSpanArgsForCall []struct {
		operationName string
		opts          []opentracing.StartSpanOption
	}
	startSpanReturns struct {
		result1 opentracing.Span
	}
	startSpanReturnsOnCall map[int]struct {
		result1 opentracing.Span
	}
	InjectStub        func(sm opentracing.SpanContext, format interface{}, carrier interface{}) error
	injectMutex       sync.RWMutex
	injectArgsForCall []struct {
		sm      opentracing.SpanContext
		format  interface{}
		carrier interface{}
	}
	injectReturns struct {
		result1 error
	}
	injectReturnsOnCall map[int]struct {
		result1 error
	}
	ExtractStub        func(format interface{}, carrier interface{}) (opentracing.SpanContext, error)
	extractMutex       sync.RWMutex
	extractArgsForCall []struct {
		format  interface{}
		carrier interface{}
	}
	extractReturns struct {
		result1 opentracing.SpanContext
		result2 error
	}
	extractReturnsOnCall map[int]struct {
		result1 opentracing.SpanContext
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTracer) StartSpan(operationName string, opts ...opentracing.StartSpanOption) opentracing.Span {
	fake.startSpanMutex.Lock()
	ret, specificReturn := fake.startSpanReturnsOnCall[len(fake.startSpanArgsForCall)]
	fake.startSpanArgsForCall = append(fake.startSpanArgsForCall, struct {
		operationName string
		opts          []opentracing.StartSpanOption
	}{operationName, opts})
	fake.recordInvocation("StartSpan", []interface{}{operationName, opts})
	fake.startSpanMutex.Unlock()
	if fake.StartSpanStub != nil {
		return fake.StartSpanStub(operationName, opts...)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.startSpanReturns.result1
}

func (fake *FakeTracer) StartSpanCallCount() int {
	fake.startSpanMutex.RLock()
	defer fake.startSpanMutex.RUnlock()
	return len(fake.startSpanArgsForCall)
}

func (fake *FakeTracer) StartSpanArgsForCall(i int) (string, []opentracing.StartSpanOption) {
	fake.startSpanMutex.RLock()
	defer fake.startSpanMutex.RUnlock()
	return fake.startSpanArgsForCall[i].operationName, fake.startSpanArgsForCall[i].opts
}

func (fake *FakeTracer) StartSpanReturns(result1 opentracing.Span) {
	fake.StartSpanStub = nil
	fake.startSpanReturns = struct {
		result1 opentracing.Span
	}{result1}
}

func (fake *FakeTracer) StartSpanReturnsOnCall(i int, result1 opentracing.Span) {
	fake.StartSpanStub = nil
	if fake.startSpanReturnsOnCall == nil {
		fake.startSpanReturnsOnCall = make(map[int]struct {
			result1 opentracing.Span
		})
	}
	fake.startSpanReturnsOnCall[i] = struct {
		result1 opentracing.Span
	}{result1}
}

func (fake *FakeTracer) Inject(sm opentracing.SpanContext, format interface{}, carrier interface{}) error {
	fake.injectMutex.Lock()
	ret, specificReturn := fake.injectReturnsOnCall[len(fake.injectArgsForCall)]
	fake.injectArgsForCall = append(fake.injectArgsForCall, struct {
		sm      opentracing.SpanContext
		format  interface{}
		carrier interface{}
	}{sm, format, carrier})
	fake.recordInvocation("Inject", []interface{}{sm, format, carrier})
	fake.injectMutex.Unlock()
	if fake.InjectStub != nil {
		return fake.InjectStub(sm, format, carrier)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.injectReturns.result1
}

func (fake *FakeTracer) InjectCallCount() int {
	fake.injectMutex.RLock()
	defer fake.injectMutex.RUnlock()
	return len(fake.injectArgsForCall)
}

func (fake *FakeTracer) InjectArgsForCall(i int) (opentracing.SpanContext, interface{}, interface{}) {
	fake.injectMutex.RLock()
	defer fake.injectMutex.RUnlock()
	return fake.injectArgsForCall[i].sm, fake.injectArgsForCall[i].format, fake.injectArgsForCall[i].carrier
}

func (fake *FakeTracer) InjectReturns(result1 error) {
	fake.InjectStub = nil
	fake.injectReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTracer) InjectReturnsOnCall(i int, result1 error) {
	fake.InjectStub = nil
	if fake.injectReturnsOnCall == nil {
		fake.injectReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.injectReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeTracer) Extract(format interface{}, carrier interface{}) (opentracing.SpanContext, error) {
	fake.extractMutex.Lock()
	ret, specificReturn := fake.extractReturnsOnCall[len(fake.extractArgsForCall)]
	fake.extractArgsForCall = append(fake.extractArgsForCall, struct {
		format  interface{}
		carrier interface{}
	}{format, carrier})
	fake.recordInvocation("Extract", []interface{}{format, carrier})
	fake.extractMutex.Unlock()
	if fake.ExtractStub != nil {
		return fake.ExtractStub(format, carrier)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.extractReturns.result1, fake.extractReturns.result2
}

func (fake *FakeTracer) ExtractCallCount() int {
	fake.extractMutex.RLock()
	defer fake.extractMutex.RUnlock()
	return len(fake.extractArgsForCall)
}

func (fake *FakeTracer) ExtractArgsForCall(i int) (interface{}, interface{}) {
	fake.extractMutex.RLock()
	defer fake.extractMutex.RUnlock()
	return fake.extractArgsForCall[i].format, fake.extractArgsForCall[i].carrier
}

func (fake *FakeTracer) ExtractReturns(result1 opentracing.SpanContext, result2 error) {
	fake.ExtractStub = nil
	fake.extractReturns = struct {
		result1 opentracing.SpanContext
		result2 error
	}{result1, result2}
}

func (fake *FakeTracer) ExtractReturnsOnCall(i int, result1 opentracing.SpanContext, result2 error) {
	fake.ExtractStub = nil
	if fake.extractReturnsOnCall == nil {
		fake.extractReturnsOnCall = make(map[int]struct {
			result1 opentracing.SpanContext
			result2 error
		})
	}
	fake.extractReturnsOnCall[i] = struct {
		result1 opentracing.SpanContext
		result2 error
	}{result1, result2}
}

func (fake *FakeTracer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.startSpanMutex.RLock()
	defer fake.startSpanMutex.RUnlock()
	fake.injectMutex.RLock()
	defer fake.injectMutex.RUnlock()
	fake.extractMutex.RLock()
	defer fake.extractMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeTracer) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ opentracing.Tracer = new(FakeTracer)
