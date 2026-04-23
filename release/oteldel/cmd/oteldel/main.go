package main

import (
	"fmt"
	"os"

	"github.com/alexfalkowski/docker/release/oteldel/internal/args"
	"github.com/alexfalkowski/docker/release/oteldel/internal/config"
	"github.com/alexfalkowski/docker/release/oteldel/internal/export"
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(argv []string, stdout, stderr *os.File) int {
	command, err := args.Parse(argv)
	if err != nil {
		fmt.Fprintf(stderr, "oteldel: %v\n", err)
		return 1
	}

	cfg, warnings := config.LoadFromEnv()
	for _, warning := range warnings {
		fmt.Fprintf(stderr, "oteldel: warning: %s\n", warning)
	}

	emitter, warning, err := export.NewEmitter(cfg)
	if warning != "" {
		fmt.Fprintf(stderr, "oteldel: warning: %s\n", warning)
	}
	if err != nil {
		fmt.Fprintf(stderr, "oteldel: warning: %v\n", err)
		return 0
	}
	defer emitter.Close()

	if err := emitter.Emit(command); err != nil {
		fmt.Fprintf(stderr, "oteldel: warning: %v\n", err)
	}

	return 0
}
