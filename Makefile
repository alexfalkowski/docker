include bin/build/make/help.mak
include bin/build/make/git.mak

# Lint all scripts.
lint:
	@scripts/lint

# Scan the repository with Trivy (CRITICAL severity).
trivy-repo:
	@$(BIN_ROOT)/build/sec/trivy-repo

# Pull latest containers, or one service with service=<name>.
pull-latest:
	@scripts/compose -f compose.yml pull $(service)

# Start all dependencies, or one service with service=<name>.
start:
	@scripts/compose -f compose.yml up -d --remove-orphans $(service)

# Stop dependencies.
stop:
	@scripts/compose -f compose.yml down --remove-orphans

# Follow logs for all dependencies, or one service with service=<name>.
logs:
	@scripts/compose -f compose.yml logs -f $(service)

# Destructively prune all unused images with Podman or Docker.
clean:
	@scripts/clean
