package export_test

import (
	"bytes"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/alexfalkowski/docker/release/oteldel/internal/args"
	"github.com/alexfalkowski/docker/release/oteldel/internal/config"
	"github.com/alexfalkowski/docker/release/oteldel/internal/export"
	"github.com/stretchr/testify/require"
	collectormetricpb "go.opentelemetry.io/proto/otlp/collector/metrics/v1"
	commonpb "go.opentelemetry.io/proto/otlp/common/v1"
	"google.golang.org/protobuf/proto"
)

func TestEmitterExportsCounterAndHistogram(t *testing.T) {
	t.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", "http://oteldel.test")

	requests := make(chan *collectormetricpb.ExportMetricsServiceRequest, 2)
	cfg := config.Config{
		HasEndpoint:     true,
		MetricsExporter: "otlp",
		MetricInterval:  60 * time.Second,
		MetricTimeout:   30 * time.Second,
		HTTPClient: &http.Client{
			Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
				body, err := io.ReadAll(r.Body)
				require.NoError(t, err)

				req := &collectormetricpb.ExportMetricsServiceRequest{}
				require.NoError(t, proto.Unmarshal(body, req))

				requests <- req

				return &http.Response{
					StatusCode: http.StatusAccepted,
					Body:       io.NopCloser(bytes.NewReader(nil)),
					Header:     make(http.Header),
				}, nil
			}),
		},
	}

	emitter, _, err := export.NewEmitter(cfg)
	require.NoError(t, err)
	defer emitter.Close()

	err = emitter.Emit(args.Request{
		MetricType: "counter",
		MetricName: "cicd.pipeline.run.errors",
		Value:      1,
		Attrs: map[string]string{
			"cicd.pipeline.name": "package",
			"error.type":         "release_tool_failed",
		},
	})
	require.NoError(t, err)

	err = emitter.Emit(args.Request{
		MetricType: "histogram",
		MetricName: "cicd.pipeline.run.duration",
		Value:      2.5,
		Attrs: map[string]string{
			"cicd.pipeline.name":      "package",
			"cicd.pipeline.run.state": "executing",
			"cicd.pipeline.result":    "success",
		},
	})
	require.NoError(t, err)

	assertMetricName(t, <-requests, "cicd.pipeline.run.errors")
	req := <-requests
	assertMetricName(t, req, "cicd.pipeline.run.duration")
	assertResourceAttr(t, req, "service.name", "release")
}

func TestEmitterHonorsServiceNameFromEnvironment(t *testing.T) {
	t.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", "http://oteldel.test")
	t.Setenv("OTEL_SERVICE_NAME", "oteldel")

	requests := make(chan *collectormetricpb.ExportMetricsServiceRequest, 1)
	cfg := config.Config{
		HasEndpoint:     true,
		MetricsExporter: "otlp",
		MetricInterval:  60 * time.Second,
		MetricTimeout:   30 * time.Second,
		HTTPClient: &http.Client{
			Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
				body, err := io.ReadAll(r.Body)
				require.NoError(t, err)

				req := &collectormetricpb.ExportMetricsServiceRequest{}
				require.NoError(t, proto.Unmarshal(body, req))

				requests <- req

				return &http.Response{
					StatusCode: http.StatusAccepted,
					Body:       io.NopCloser(bytes.NewReader(nil)),
					Header:     make(http.Header),
				}, nil
			}),
		},
	}

	emitter, _, err := export.NewEmitter(cfg)
	require.NoError(t, err)
	defer emitter.Close()

	err = emitter.Emit(args.Request{
		MetricType: "counter",
		MetricName: "cicd.pipeline.run.errors",
		Value:      1,
		Attrs: map[string]string{
			"cicd.pipeline.name": "package",
			"error.type":         "release_tool_failed",
		},
	})
	require.NoError(t, err)

	assertResourceAttr(t, <-requests, "service.name", "oteldel")
}

func assertMetricName(t *testing.T, req *collectormetricpb.ExportMetricsServiceRequest, want string) {
	t.Helper()

	for _, rm := range req.ResourceMetrics {
		for _, sm := range rm.ScopeMetrics {
			for _, metric := range sm.Metrics {
				if metric.Name == want {
					return
				}
			}
		}
	}

	require.FailNow(t, "expected metric in request", "metric=%s", want)
}

func assertResourceAttr(t *testing.T, req *collectormetricpb.ExportMetricsServiceRequest, key, want string) {
	t.Helper()

	for _, rm := range req.ResourceMetrics {
		for _, attr := range rm.Resource.Attributes {
			if attr.Key == key && attr.GetValue().GetStringValue() == want {
				return
			}
		}
	}

	got := make([]string, 0)
	for _, rm := range req.ResourceMetrics {
		for _, attr := range rm.Resource.Attributes {
			got = append(got, formatAttr(attr))
		}
	}

	require.FailNow(t, "expected resource attribute in request", "attribute=%s want=%q got=%v", key, want, got)
}

func formatAttr(attr *commonpb.KeyValue) string {
	return attr.Key + "=" + attr.GetValue().GetStringValue()
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (fn roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return fn(r)
}
