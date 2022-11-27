package opentelemetry

import (
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"testing"
)

func TestBase(t *testing.T) {
	propagator, _ := otel.GetTextMapPropagator(), make(propagation.MapCarrier)

	for _, f := range propagator.Fields() {
		fmt.Println(f)
	}
}
