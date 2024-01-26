package tracer

import (
	"context"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

type TracerWithoutExport struct{}

func (ce *TracerWithoutExport) ExportSpans(ctx context.Context, spans []sdktrace.ReadOnlySpan) error {
	return nil
}

func (ce *TracerWithoutExport) Shutdown(ctx context.Context) error {
	return nil
}
