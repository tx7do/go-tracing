package tracing

import (
	"context"
)

type noop struct{}

func NewTracer(_ ...Option) Tracer {
	return &noop{}
}

func (n noop) Init(_ ...Option) error {
	return nil
}

func (n noop) Start(_ context.Context, _ string) (context.Context, *Span) {
	return nil, nil
}

func (n noop) Finish(_ *Span) error {
	return nil
}

func (n noop) Read(_ ...ReadOption) ([]*Span, error) {
	return nil, nil
}
