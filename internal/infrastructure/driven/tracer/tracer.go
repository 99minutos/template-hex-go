package tracer

import (
	"context"
	"fmt"
	"service/internal/infrastructure/driven/core"
	"sync"

	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	gcppropagator "github.com/GoogleCloudPlatform/opentelemetry-operations-go/propagator"
	"go.opentelemetry.io/contrib/detectors/gcp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/api/option"
)

var (
	tracerInstance trace.Tracer
	once           sync.Once
)

func GetTracer() trace.Tracer {
	once.Do(func() {
		tracerInstance = NewTracer()
	})

	return tracerInstance
}

func NewTracer() trace.Tracer {
	ctx := context.Background()
	config := core.GetEnviroments()

	exporter, err := getExporter(config.ProjectId)
	if err != nil {
		println("failed creating tracer", "error", err)
	}

	res, err := newResource(ctx, config.AppName)
	if err != nil {
		println("failed creating tracer", "error", err)
	}

	provider := sdktrace.NewTracerProvider(
		sdktrace.WithResource(res),
		sdktrace.WithBatcher(exporter),
	)
	tracer := otel.Tracer(config.AppName)
	defer func(provider *sdktrace.TracerProvider, ctx context.Context) {
		err := provider.ForceFlush(ctx)
		if err != nil {
			println("provider.ForceFlush: %v", err)
		}
	}(provider, ctx)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			gcppropagator.CloudTraceOneWayPropagator{},
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)
	otel.SetTracerProvider(provider)
	return tracer
}

func getExporter(projectId string) (sdktrace.SpanExporter, error) {
	gcpExporter, err := texporter.New(
		texporter.WithProjectID(projectId),
		texporter.WithTraceClientOptions(
			[]option.ClientOption{option.WithTelemetryDisabled()},
		),
	)
	if err != nil {
		stdoutExporter, err := stdouttrace.New()
		if err != nil {
			return nil, err
		}
		return stdoutExporter, nil
	}

	return gcpExporter, nil
}

func newResource(ctx context.Context, appName string) (*resource.Resource, error) {
	r, err := resource.New(ctx,
		// Use the GCP resource detector!
		resource.WithDetectors(gcp.NewDetector()),
		// Keep the default detectors
		resource.WithTelemetrySDK(),
		// Add your own custom attributes to identify your application
		resource.WithAttributes(
			semconv.ServiceNameKey.String(appName),
		),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create resource for telemetry: %w", err)
	}
	return r, nil
}
