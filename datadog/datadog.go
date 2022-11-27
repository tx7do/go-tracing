package datadog

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-tracing"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func StartSpanFromContext(ctx context.Context, operationName string, opts ...tracer.StartSpanOption) (tracer.Span, context.Context) {
	md, ok := tracing.FromContext(ctx)
	if !ok {
		md = make(map[string]string)
	}

	if spanCtx, err := tracer.Extract(tracer.TextMapCarrier(md)); err == nil {
		opts = append(opts, tracer.ChildOf(spanCtx))
	}

	span, ctx := tracer.StartSpanFromContext(ctx, operationName, opts...)

	if err := tracer.Inject(span.Context(), tracer.TextMapCarrier(md)); err != nil {
		log.Errorf("error while injecting trace to context: %s\n", err)
	}

	return span, tracing.NewContext(ctx, md)
}
