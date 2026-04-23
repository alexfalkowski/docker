// Package main provides the oteldel command-line entrypoint.
//
// The oteldel binary is an internal release-tooling command that emits OTLP
// metrics for delivery events such as pipeline runs, artifact publication, and
// deployment requests. It supports a small command surface:
//
//	emit <counter|histogram> <metric> <value> [--attr key=value...]
//
// The command is intentionally best-effort. Exporter construction and metric
// delivery failures are reported to stderr without failing the surrounding
// release flow.
package main
