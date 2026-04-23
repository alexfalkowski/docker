package export

import (
	"context"
	"fmt"
	"time"

	"github.com/alexfalkowski/docker/release/oteldel/internal/args"
	"github.com/alexfalkowski/docker/release/oteldel/internal/config"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
)

const scopeName = "github.com/alexfalkowski/docker/release/oteldel"
const defaultServiceName = "release"

// Emitter records oteldel commands as metrics.
type Emitter interface {
	// Emit records a single parsed metric command.
	Emit(args.Request) error
	// Close releases exporter resources.
	Close() error
}

type noopEmitter struct{}

func (noopEmitter) Emit(args.Request) error {
	return nil
}

func (noopEmitter) Close() error {
	return nil
}

type metricEmitter struct {
	provider *sdkmetric.MeterProvider
	meter    metric.Meter
	timeout  time.Duration
}

// NewEmitter constructs an Emitter from the supplied configuration.
//
// When telemetry is disabled, no metrics exporter is configured, or no OTLP
// endpoint is available, it returns a no-op emitter so callers can continue
// without special-case branching.
func NewEmitter(cfg config.Config) (Emitter, string, error) {
	if cfg.Disabled || cfg.MetricsExporter == "none" || !cfg.HasEndpoint {
		return noopEmitter{}, "", nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.MetricTimeout)
	defer cancel()

	options := []otlpmetrichttp.Option{}
	if cfg.HTTPClient != nil {
		options = append(options, otlpmetrichttp.WithHTTPClient(cfg.HTTPClient))
	}

	exporter, err := otlpmetrichttp.New(ctx, options...)
	if err != nil {
		return noopEmitter{}, "", fmt.Errorf("create metrics exporter: %w", err)
	}

	res, err := resource.New(
		ctx,
		resource.WithAttributes(semconv.ServiceName(defaultServiceName)),
		resource.WithFromEnv(),
	)
	if err != nil {
		return noopEmitter{}, "", fmt.Errorf("create resource: %w", err)
	}

	reader := sdkmetric.NewPeriodicReader(
		exporter,
		sdkmetric.WithInterval(cfg.MetricInterval),
		sdkmetric.WithTimeout(cfg.MetricTimeout),
	)

	provider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(reader),
		sdkmetric.WithResource(res),
	)

	return &metricEmitter{
		provider: provider,
		meter:    provider.Meter(scopeName),
		timeout:  cfg.MetricTimeout,
	}, "", nil
}

func (e *metricEmitter) Emit(command args.Request) error {
	ctx, cancel := context.WithTimeout(context.Background(), e.timeout)
	defer cancel()

	attrs := make([]attribute.KeyValue, 0, len(command.Attrs))
	for key, value := range command.Attrs {
		attrs = append(attrs, attribute.String(key, value))
	}

	switch command.MetricType {
	case "counter":
		counter, err := e.meter.Float64Counter(command.MetricName)
		if err != nil {
			return fmt.Errorf("create counter %q: %w", command.MetricName, err)
		}
		counter.Add(ctx, command.Value, metric.WithAttributes(attrs...))
	case "histogram":
		histogram, err := e.meter.Float64Histogram(command.MetricName)
		if err != nil {
			return fmt.Errorf("create histogram %q: %w", command.MetricName, err)
		}
		histogram.Record(ctx, command.Value, metric.WithAttributes(attrs...))
	default:
		return fmt.Errorf("unsupported metric type %q", command.MetricType)
	}

	if err := e.provider.ForceFlush(ctx); err != nil {
		return fmt.Errorf("flush metrics: %w", err)
	}

	return nil
}

func (e *metricEmitter) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), e.timeout)
	defer cancel()

	return e.provider.Shutdown(ctx)
}
