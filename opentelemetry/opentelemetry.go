package opentelemetry

import (
	"context"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"

	"github.com/tx7do/go-tracing"
)

const (
	instrumentationName = "trace/opentelemetry"
)

// StartSpanFromContext returns a new span with the given operation name and options. If a span
// is found in the context, it will be used as the parent of the resulting span.
func StartSpanFromContext(ctx context.Context, tp trace.TracerProvider, name string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	md, ok := tracing.FromContext(ctx)
	if !ok {
		md = make(tracing.TraceData)
	}
	propagator, carrier := otel.GetTextMapPropagator(), make(propagation.MapCarrier)
	for k, v := range md {
		for _, f := range propagator.Fields() {
			if strings.EqualFold(k, f) {
				carrier[f] = v
			}
		}
	}
	ctx = propagator.Extract(ctx, carrier)
	spanCtx := trace.SpanContextFromContext(ctx)
	ctx = baggage.ContextWithBaggage(ctx, baggage.FromContext(ctx))

	var tracer trace.Tracer
	var span trace.Span
	if tp != nil {
		tracer = tp.Tracer(instrumentationName)
	} else {
		tracer = otel.Tracer(instrumentationName)
	}
	ctx, span = tracer.Start(trace.ContextWithRemoteSpanContext(ctx, spanCtx), name, opts...)

	carrier = make(propagation.MapCarrier)
	propagator.Inject(ctx, carrier)

	for k, v := range carrier {
		md.Set(cases.Title(language.AmericanEnglish).String(k), v)
	}
	ctx = tracing.NewContext(ctx, md)

	return ctx, span
}
