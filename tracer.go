package tracing

import (
	"context"
	"time"
)

var (
	// DefaultTracer is the default tracer.
	DefaultTracer = NewTracer()
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
