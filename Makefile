include bin/build/make/help.mak
include bin/build/make/git.mak

# Lint all scripts.
lint:
	@scripts/lint

# Scan the repository with Trivy (CRITICAL severity).
trivy-repo:
	@$(BIN_ROOT)/build/sec/trivy-repo

# Pull latest containers.
pull-latest:
	@scripts/compose -f compose.yml pull $(service)

# Start dependencies.
start:
	@scripts/compose -f compose.yml up -d --remove-orphans $(service)

# Stop dependencies.
stop:
	@scripts/compose -f compose.yml down --remove-orphans

# Logs from a service.
logs:
	@scripts/compose -f compose.yml logs -f $(service)

# Clean unused Podman or Docker images.
clean:
	@scripts/clean
