package opentracing

import (
	"context"

	"github.com/opentracing/opentracing-go"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/tx7do/go-tracing"
)

func StartSpanFromContext(ctx context.Context, tracer opentracing.Tracer, spanName string, opts ...opentracing.StartSpanOption) (context.Context, opentracing.Span, error) {
	md, ok := tracing.FromContext(ctx)
	if !ok {
		md = make(tracing.TraceData)
	}

	if parentSpan := opentracing.SpanFromContext(ctx); parentSpan != nil {
		opts = append(opts, opentracing.ChildOf(parentSpan.Context()))
	} else if spanCtx, err := tracer.Extract(opentracing.TextMap, opentracing.TextMapCarrier(md)); err == nil {
		opts = append(opts, opentracing.ChildOf(spanCtx))
	}

	nmd := make(tracing.TraceData, 1)

	sp := tracer.StartSpan(spanName, opts...)

	if err := sp.Tracer().Inject(sp.Context(), opentracing.TextMap, opentracing.TextMapCarrier(nmd)); err != nil {
		return nil, nil, err
	}

	for k, v := range nmd {
		md.Set(cases.Title(language.AmericanEnglish).String(k), v)
	}

	ctx = opentracing.ContextWithSpan(ctx, sp)
	ctx = tracing.NewContext(ctx, md)
	return ctx, sp, nil
}
