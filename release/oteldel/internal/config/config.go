package config

import (
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	defaultMetricExportInterval = 60 * time.Second
	defaultMetricExportTimeout  = 30 * time.Second
)

// Config describes the oteldel runtime configuration.
type Config struct {
	// Disabled reports whether telemetry has been disabled explicitly.
	Disabled bool
	// HasEndpoint reports whether an OTLP metrics endpoint has been configured.
	HasEndpoint bool
	// MetricsExporter is the selected metrics exporter kind, such as otlp or none.
	MetricsExporter string
	// MetricInterval is the periodic export interval used by the metrics reader.
	MetricInterval time.Duration
	// MetricTimeout is the timeout applied to export and flush operations.
	MetricTimeout time.Duration
	// HTTPClient overrides the exporter transport for tests and custom callers.
	HTTPClient *http.Client
}

// LoadFromEnv loads Config from the supported OpenTelemetry environment
// variables.
//
// It returns the parsed configuration plus any non-fatal warnings for ignored
// or unsupported values.
func LoadFromEnv() (Config, []string) {
	cfg := Config{
		MetricsExporter: "otlp",
		MetricInterval:  defaultMetricExportInterval,
		MetricTimeout:   defaultMetricExportTimeout,
	}
	warnings := []string{}

	cfg.Disabled = parseBoolEnv("OTEL_SDK_DISABLED", false, &warnings)
	cfg.HasEndpoint = hasEndpoint()
	cfg.MetricsExporter = parseMetricsExporter(&warnings)
	cfg.MetricTimeout = parseDurationMillisEnv(
		"OTEL_METRIC_EXPORT_TIMEOUT",
		defaultMetricExportTimeout,
		&warnings,
	)
	cfg.MetricInterval = parseDurationMillisEnv(
		"OTEL_METRIC_EXPORT_INTERVAL",
		defaultMetricExportInterval,
		&warnings,
	)

	return cfg, warnings
}

func parseMetricsExporter(warnings *[]string) string {
	exporter := strings.TrimSpace(os.Getenv("OTEL_METRICS_EXPORTER"))
	if exporter == "" {
		return "otlp"
	}

	switch strings.ToLower(exporter) {
	case "otlp":
		return "otlp"
	case "none":
		return "none"
	default:
		*warnings = append(*warnings, "ignoring unsupported OTEL_METRICS_EXPORTER value "+strconv.Quote(exporter))
		return "otlp"
	}
}

func parseDurationMillisEnv(name string, fallback time.Duration, warnings *[]string) time.Duration {
	value := strings.TrimSpace(os.Getenv(name))
	if value == "" {
		return fallback
	}

	raw, err := strconv.Atoi(value)
	if err != nil || raw < 0 {
		*warnings = append(*warnings, "ignoring invalid "+name+" value "+strconv.Quote(value))
		return fallback
	}

	return time.Duration(raw) * time.Millisecond
}

func parseBoolEnv(name string, fallback bool, warnings *[]string) bool {
	return parseBoolValue(name, strings.TrimSpace(os.Getenv(name)), fallback, warnings)
}

func parseBoolValue(name, value string, fallback bool, warnings *[]string) bool {
	if value == "" {
		return fallback
	}

	switch strings.ToLower(value) {
	case "true":
		return true
	case "false":
		return false
	default:
		*warnings = append(*warnings, "ignoring invalid "+name+" value "+strconv.Quote(value))
		return fallback
	}
}

func hasEndpoint() bool {
	return firstNonEmptyEnv(
		"OTEL_EXPORTER_OTLP_METRICS_ENDPOINT",
		"OTEL_EXPORTER_OTLP_ENDPOINT",
	) != ""
}

func firstNonEmptyEnv(names ...string) string {
	for _, name := range names {
		if value := strings.TrimSpace(os.Getenv(name)); value != "" {
			return value
		}
	}

	return ""
}
