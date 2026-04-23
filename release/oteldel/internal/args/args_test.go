package args_test

import (
	"testing"

	"github.com/alexfalkowski/docker/release/oteldel/internal/args"
	"github.com/stretchr/testify/require"
)

func TestParseEmit(t *testing.T) {
	command, err := args.Parse([]string{
		"emit",
		"counter",
		"cicd.pipeline.run.errors",
		"1",
		"--attr",
		"cicd.pipeline.name=package",
		"-a",
		"error.type=release_tool_failed",
	})
	require.NoError(t, err)
	require.Equal(t, "counter", command.MetricType)
	require.Equal(t, "package", command.Attrs["cicd.pipeline.name"])
}

func TestParseRejectsInvalidAttribute(t *testing.T) {
	_, err := args.Parse([]string{
		"emit",
		"histogram",
		"cicd.pipeline.run.duration",
		"1.5",
		"--attr",
		"bad-attr",
	})
	require.Error(t, err)
}

func TestParseRejectsUnknownCommand(t *testing.T) {
	_, err := args.Parse([]string{"nope"})
	require.Error(t, err)
}
