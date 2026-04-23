// Package config loads the oteldel runtime configuration from environment
// variables.
//
// The package reads the small subset of OpenTelemetry environment variables
// that oteldel needs to decide whether metric export should run and how often
// the exporter should flush.
package config
