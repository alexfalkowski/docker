include bin/build/make/git.mak

# Lint all scripts.
scripts-lint:
	shellcheck scripts/*.sh scripts/build scripts/clean scripts/compose scripts/lint scripts/push

# Lint all the images.
docker-lint:
	scripts/lint

# Run all linters.
lint: docker-lint scripts-lint

# Build all the images.
docker-build:
	scripts/build

# Push all the images.
docker-push:
	scripts/push

# Pull latest containers.
docker-pull:
	scripts/compose -f compose.yml pull $(service)

# Start dependencies.
start:
	scripts/compose -f compose.yml up -d --remove-orphans $(service)

# Stop dependencies.
stop:
	scripts/compose -f compose.yml down --remove-orphans

# Logs from a service.
logs:
	scripts/compose -f compose.yml logs -f $(service)

# Clean all unused docker images.
clean:
	scripts/clean image prune -a -f
