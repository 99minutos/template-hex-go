package tracer

import (
	"context"
	"example-service/internal/infraestructure/driven/core"
	"fmt"
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
	"go.uber.org/zap"
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
	defaultLog := core.GetDefaultLogger()
	defaultLog.Infow("Tracer is starting...")
	config := core.GetEnviroments()

	exporter, err := getExporter(config.ProjectId, defaultLog)
	if err != nil {
		defaultLog.Errorw("failed creating tracer", "error", err)
	}

	res, err := newResource(ctx, config.AppName)
	if err != nil {
		defaultLog.Errorw("failed creating tracer", "error", err)
	}

	provider := sdktrace.NewTracerProvider(
		sdktrace.WithResource(res),
		sdktrace.WithBatcher(exporter),
	)
	tracer := otel.Tracer(config.AppName)
	defaultLog.Infow("Tracer has started.")
	defer func(provider *sdktrace.TracerProvider, ctx context.Context) {
		err := provider.ForceFlush(ctx)
		if err != nil {
			defaultLog.Warnw("provider.ForceFlush: %v", err)
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

func getExporter(projectId string, log zap.SugaredLogger) (sdktrace.SpanExporter, error) {
	gcpExporter, err := texporter.New(texporter.WithProjectID(projectId))
	if err != nil {
		log.Warnw("Failed to create the Google Cloud Trace exporter, using console exporter instead", "err", err)
		stdoutExporter, err := stdouttrace.New()
		if err != nil {
			log.Errorw("Failed to create the console exporter", "err", err)
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
