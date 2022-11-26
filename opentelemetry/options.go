package opentelemetry

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport"
	"github.com/tx7do/kratos-transport/broker"
	"go.opentelemetry.io/otel/trace"
)

type Options struct {
	TraceProvider trace.TracerProvider

	CallFilter       CallFilter
	StreamFilter     StreamFilter
	PublishFilter    PublishFilter
	SubscriberFilter SubscriberFilter
	HandlerFilter    HandlerFilter
}

// CallFilter used to filter client.Call, return true to skip call trace.
type CallFilter func(context.Context, transport.Header) bool

// StreamFilter used to filter client.Stream, return true to skip stream trace.
type StreamFilter func(context.Context, transport.Header) bool

// PublishFilter used to filter client.Publish, return true to skip publish trace.
type PublishFilter func(context.Context, broker.Message) bool

// SubscriberFilter used to filter server.Subscribe, return true to skip subscribe trace.
type SubscriberFilter func(context.Context, broker.Message) bool

// HandlerFilter used to filter server.Handle, return true to skip handle trace.
type HandlerFilter func(context.Context, transport.Header) bool

type Option func(*Options)

func WithTraceProvider(tp trace.TracerProvider) Option {
	return func(o *Options) {
		o.TraceProvider = tp
	}
}

func WithCallFilter(filter CallFilter) Option {
	return func(o *Options) {
		o.CallFilter = filter
	}
}

func WithStreamFilter(filter StreamFilter) Option {
	return func(o *Options) {
		o.StreamFilter = filter
	}
}

func WithPublishFilter(filter PublishFilter) Option {
	return func(o *Options) {
		o.PublishFilter = filter
	}
}

func WithSubscribeFilter(filter SubscriberFilter) Option {
	return func(o *Options) {
		o.SubscriberFilter = filter
	}
}

func WithHandleFilter(filter HandlerFilter) Option {
	return func(o *Options) {
		o.HandlerFilter = filter
	}
}
