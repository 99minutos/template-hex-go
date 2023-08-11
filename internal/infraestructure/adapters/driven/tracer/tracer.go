package tracer

import (
	"context"
	"example-service/internal/config"
	"example-service/internal/infraestructure/adapters/driven/logger"
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

func Setup(ctx context.Context, cfg *config.AppConfig) (trace.Tracer, error) {
	logger := logger.Logger
	logger.Info("Tracer is starting...")

	exporter, err := getExporter(ctx, cfg)
	if err != nil {
		logger.Warnf("failed creating tracer exporter: %v", err.Error())
	}

	res, err := newResource(ctx, cfg)
	if err != nil {
		logger.Warnf("failed creating tracer resource: %v", err.Error())
	}

	provider := sdktrace.NewTracerProvider(
		sdktrace.WithResource(res),
		sdktrace.WithBatcher(exporter),
	)

	TelemetryProvider = provider
	TelemetryShutdown = provider.Shutdown
	Tracer = otel.Tracer("github.com/99minutos/inbound-service")
	logger.Info("Tracer has started.")
	defer func(provider *sdktrace.TracerProvider, ctx context.Context) {
		err := provider.ForceFlush(ctx)
		if err != nil {
			logger.Warnf("provider.ForceFlush: %v", err)
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

func getExporter(ctx context.Context, cfg *config.AppConfig) (sdktrace.SpanExporter, error) {
	logger := logger.Logger
	exporter, err := texporter.New(texporter.WithProjectID(cfg.ProjectId))
	if err != nil {
		logger.Warnf("texporter.New: %v", err)
	}

	return exporter, err
}

func newResource(ctx context.Context, cfg *config.AppConfig) (*resource.Resource, error) {
	r, err := resource.New(ctx,
		// Use the GCP resource detector!
		resource.WithDetectors(gcp.NewDetector()),
		// Keep the default detectors
		resource.WithTelemetrySDK(),
		// Add your own custom attributes to identify your application
		resource.WithAttributes(
			semconv.ServiceNameKey.String(cfg.AppName),
		),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create resource for telemetry: %w", err)
	}
	return r, nil
}
