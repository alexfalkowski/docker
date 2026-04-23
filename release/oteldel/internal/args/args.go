package args

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Request describes a parsed oteldel invocation.
type Request struct {
	// MetricType is the metric instrument kind to emit, such as counter or histogram.
	MetricType string
	// MetricName is the fully-qualified OpenTelemetry metric name.
	MetricName string
	// Value is the numeric observation recorded for the metric.
	Value float64
	// Attrs holds the string attributes supplied on the command line.
	Attrs map[string]string
}

// Parse converts command-line arguments into a Request.
//
// It supports the "emit" command with a metric type, metric name, numeric
// value, and optional repeated attribute flags.
func Parse(args []string) (Request, error) {
	if len(args) == 0 {
		return Request{}, errors.New("missing command")
	}

	switch args[0] {
	case "emit":
		return parseEmit(args[1:])
	default:
		return Request{}, fmt.Errorf("unknown command %q", args[0])
	}
}

func parseEmit(args []string) (Request, error) {
	if len(args) < 3 {
		return Request{}, errors.New("emit requires <counter|histogram> <metric> <value>")
	}

	metricType := args[0]
	if metricType != "counter" && metricType != "histogram" {
		return Request{}, fmt.Errorf("unsupported metric type %q", metricType)
	}

	value, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		return Request{}, fmt.Errorf("invalid metric value %q: %w", args[2], err)
	}

	attrs, err := parseAttrs(args[3:])
	if err != nil {
		return Request{}, err
	}

	return Request{
		MetricType: metricType,
		MetricName: args[1],
		Value:      value,
		Attrs:      attrs,
	}, nil
}

func parseAttrs(args []string) (map[string]string, error) {
	attrs := map[string]string{}

	for idx := 0; idx < len(args); idx++ {
		switch args[idx] {
		case "--attr", "-a":
			idx++
			if idx >= len(args) {
				return nil, fmt.Errorf("missing value for %s", args[idx-1])
			}

			key, value, ok := strings.Cut(args[idx], "=")
			if !ok || key == "" {
				return nil, fmt.Errorf("invalid attribute %q", args[idx])
			}

			attrs[key] = value
		default:
			return nil, fmt.Errorf("unknown argument %q", args[idx])
		}
	}

	return attrs, nil
}
