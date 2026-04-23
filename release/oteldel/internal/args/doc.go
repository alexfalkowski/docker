// Package args parses oteldel command-line arguments into typed commands.
//
// It is responsible only for syntactic validation and conversion of command
// line input into a structured Request value. Semantic handling of the parsed
// command, including configuration and metric emission, is delegated to other
// packages.
package args
