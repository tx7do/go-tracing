package tracing

import (
	"context"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	TracerName = "Kratos-Tracer"
	SpanID     = "Kratos-Span-ID"
	TraceIDKey = "Kratos-Trace-ID"
)

type traceDataKey struct{}

type TraceData map[string]string

func (td TraceData) Get(key string) (string, bool) {
	// attempt to get as is
	val, ok := td[key]
	if ok {
		return val, ok
	}

	// attempt to get lower case
	val, ok = td[cases.Title(language.AmericanEnglish).String(key)]
	return val, ok
}

func (td TraceData) Set(key, val string) {
	td[key] = val
}

func (td TraceData) Delete(key string) {
	// delete key as-is
	delete(td, key)
	// delete also Title key
	delete(td, cases.Title(language.AmericanEnglish).String(key))
}

// Copy makes a copy of the TraceData
func Copy(td TraceData) TraceData {
	cmd := make(TraceData, len(td))
	for k, v := range td {
		cmd[k] = v
	}
	return cmd
}

// Delete key from TraceData
func Delete(ctx context.Context, k string) context.Context {
	return Set(ctx, k, "")
}

// Set add key with val to TraceData
func Set(ctx context.Context, k, v string) context.Context {
	md, ok := FromContext(ctx)
	if !ok {
		md = make(TraceData)
	}
	if v == "" {
		delete(md, k)
	} else {
		md[k] = v
	}
	return context.WithValue(ctx, traceDataKey{}, md)
}

// Get returns a single value from TraceData in the context
func Get(ctx context.Context, key string) (string, bool) {
	md, ok := FromContext(ctx)
	if !ok {
		return "", ok
	}
	// attempt to get as is
	val, ok := md[key]
	if ok {
		return val, ok
	}

	// attempt to get lower case
	val, ok = md[cases.Title(language.AmericanEnglish).String(key)]

	return val, ok
}

// FromContext returns TraceData from the given context
func FromContext(ctx context.Context) (TraceData, bool) {
	md, ok := ctx.Value(traceDataKey{}).(TraceData)
	if !ok {
		return nil, ok
	}

	// capitalise all values
	newMD := make(TraceData, len(md))
	for k, v := range md {
		newMD[cases.Title(language.AmericanEnglish).String(k)] = v
	}

	return newMD, ok
}

// NewContext creates a new context with the given TraceData
func NewContext(ctx context.Context, md TraceData) context.Context {
	return context.WithValue(ctx, traceDataKey{}, md)
}

// MergeContext merges TraceData to existing TraceData, overwriting if specified
func MergeContext(ctx context.Context, patchMd TraceData, overwrite bool) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	md, _ := ctx.Value(traceDataKey{}).(TraceData)
	cmd := make(TraceData, len(md))
	for k, v := range md {
		cmd[k] = v
	}
	for k, v := range patchMd {
		if _, ok := cmd[k]; ok && !overwrite {
			// skip
		} else if v != "" {
			cmd[k] = v
		} else {
			delete(cmd, k)
		}
	}
	return context.WithValue(ctx, traceDataKey{}, cmd)
}
