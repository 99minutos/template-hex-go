package tracer

import (
	"context"
	"example-service/internal/domain/core"
	"fmt"
	"go.opentelemetry.io/contrib/detectors/gcp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"

	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	gcppropagator "github.com/GoogleCloudPlatform/opentelemetry-operations-go/propagator"
)

var TelemetryProvider *sdktrace.TracerProvider
var Tracer trace.Tracer
var TelemetryShutdown func(ctx context.Context) error

func GetTracerInstance() trace.Tracer {
	return Tracer
}

func Setup(ctx context.Context, acx *core.AppContext) (trace.Tracer, error) {
	acx.Infow("Tracer is starting...")

	exporter, err := getExporter(ctx, acx)
	if err != nil {
		acx.Warnw("failed creating tracer exporter", "error", err)
		exporter = &TracerWithoutExport{}
	}

	res, err := newResource(ctx, acx)
	if err != nil {
		acx.Warnw("failed creating tracer resource", "error", err)
	}

	provider := sdktrace.NewTracerProvider(
		sdktrace.WithResource(res),
		sdktrace.WithBatcher(exporter),
	)

	TelemetryProvider = provider
	TelemetryShutdown = provider.Shutdown
	Tracer = otel.Tracer(acx.Envs.AppName)
	acx.Infow("Tracer has started.")
	defer func(provider *sdktrace.TracerProvider, ctx context.Context) {
		err := provider.ForceFlush(ctx)
		if err != nil {
			acx.Warnw("provider.ForceFlush: %v", err)
		}
	}(provider, ctx)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			gcppropagator.CloudTraceOneWayPropagator{},
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)
	otel.SetTracerProvider(TelemetryProvider)
	return Tracer, nil
}

func getExporter(ctx context.Context, acx *core.AppContext) (sdktrace.SpanExporter, error) {
	exporter, err := texporter.New(texporter.WithProjectID(acx.Envs.ProjectId))
	if err != nil {
		acx.Warnw("failed creating tracer exporter", "error", err)
		return nil, err
	}

	return exporter, nil
}

func newResource(ctx context.Context, acx *core.AppContext) (*resource.Resource, error) {
	r, err := resource.New(ctx,
		// Use the GCP resource detector!
		resource.WithDetectors(gcp.NewDetector()),
		// Keep the default detectors
		resource.WithTelemetrySDK(),
		// Add your own custom attributes to identify your application
		resource.WithAttributes(
			semconv.ServiceNameKey.String(acx.Envs.AppName),
		),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create resource for telemetry: %w", err)
	}
	return r, nil
}
