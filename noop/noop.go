package noop

import (
	"context"
	"github.com/tx7do/go-tracing"
)

type noop struct{}

func NewTracer(_ ...tracing.Option) tracing.Tracer {
	return &noop{}
}

func (n noop) Init(_ ...tracing.Option) error {
	return nil
}

func (n noop) Start(_ context.Context, _ string) (context.Context, *tracing.Span) {
	return nil, nil
}

func (n noop) Finish(_ *tracing.Span) error {
	return nil
}

func (n noop) Read(_ ...tracing.ReadOption) ([]*tracing.Span, error) {
	return nil, nil
}
