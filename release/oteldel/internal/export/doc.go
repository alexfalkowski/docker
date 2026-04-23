// Package export emits oteldel metrics through OpenTelemetry exporters.
//
// It translates parsed CLI commands into OpenTelemetry metric instruments,
// configures an OTLP HTTP exporter, and applies the release tool's default
// resource identity while still honoring standard environment-based resource
// overrides.
package export
