package opencensus

import (
	"context"
	"encoding/base64"

	"go.opencensus.io/trace"
	"go.opencensus.io/trace/propagation"

	"github.com/tx7do/go-tracing"
)

const (
	TracePropagationField = "X-Trace-Context"
)

func injectTraceIntoCtx(ctx context.Context, span *trace.Span) context.Context {
	spanCtx := propagation.Binary(span.SpanContext())
	return tracing.Set(ctx, TracePropagationField, base64.RawStdEncoding.EncodeToString(spanCtx))
}
