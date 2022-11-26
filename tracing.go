package tracing

import (
	"context"
	"time"

	"github.com/tx7do/go-tracing/noop"
)

var (
	// DefaultTracer is the default tracer.
	DefaultTracer = noop.NewTracer()
)

// Tracer is an interface for distributed tracing.
type Tracer interface {
	// Start a trace
	Start(ctx context.Context, name string) (context.Context, *Span)
	// Finish the trace
	Finish(*Span) error
	// Read the traces
	Read(...ReadOption) ([]*Span, error)
}

// SpanType describe the nature of the trace span.
type SpanType int

const (
	// SpanTypeRequestInbound is a span created when serving a request.
	SpanTypeRequestInbound SpanType = iota
	// SpanTypeRequestOutbound is a span created when making a service call.
	SpanTypeRequestOutbound
)

// Span is used to record an entry.
type Span struct {
	// id of the Trace
	Trace string
	// Name of the span
	Name string
	// Id .
	Id string
	// Type .
	Type SpanType
	// Parent span id
	Parent string
	// Started time
	Started time.Time
	// Duration in nano seconds
	Duration time.Duration
	// Metadata associated data
	Metadata map[string]string
}

// TraceIdFromContext returns a span from context.
func TraceIdFromContext(ctx context.Context) (traceID string, parentSpanID string, isFound bool) {
	traceID, traceOk := Get(ctx, TraceIDKey)

	if !traceOk {
		isFound = false
		return
	}

	if !traceOk {
		traceID = TracerName
	}

	parentSpanID, ok := Get(ctx, SpanID)

	return traceID, parentSpanID, ok
}

// TraceIdToContext saves the trace and span ids in the context.
func TraceIdToContext(ctx context.Context, traceID, parentSpanID string) context.Context {
	return MergeContext(ctx, map[string]string{
		TraceIDKey: traceID,
		SpanID:     parentSpanID,
	}, true)
}
