package config_test

import (
	"testing"
	"time"

	"github.com/alexfalkowski/docker/release/oteldel/internal/config"
	"github.com/stretchr/testify/require"
)

func TestLoadFromEnvDefaults(t *testing.T) {
	cfg, warnings := config.LoadFromEnv()
	require.Empty(t, warnings)
	require.False(t, cfg.HasEndpoint)
	require.Equal(t, "otlp", cfg.MetricsExporter)
}

func TestLoadFromEnvDetectsEndpoint(t *testing.T) {
	t.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", "http://127.0.0.1:4318")

	cfg, warnings := config.LoadFromEnv()
	require.Empty(t, warnings)
	require.True(t, cfg.HasEndpoint)
}

func TestLoadFromEnvNoEndpointIsAllowed(t *testing.T) {
	cfg, _ := config.LoadFromEnv()
	require.False(t, cfg.HasEndpoint)
}

func TestLoadFromEnvParsesExporterConfig(t *testing.T) {
	t.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", "http://127.0.0.1:4318")
	t.Setenv("OTEL_METRIC_EXPORT_TIMEOUT", "1500")

	cfg, warnings := config.LoadFromEnv()
	require.Empty(t, warnings)
	require.True(t, cfg.HasEndpoint)
	require.Equal(t, 1500*time.Millisecond, cfg.MetricTimeout)
}

func TestLoadFromEnvSupportsNoMetricsExporter(t *testing.T) {
	t.Setenv("OTEL_METRICS_EXPORTER", "none")

	cfg, warnings := config.LoadFromEnv()
	require.Empty(t, warnings)
	require.Equal(t, "none", cfg.MetricsExporter)
}
