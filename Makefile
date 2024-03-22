# Lint all the images.
lint:
	scripts/lint

# Build all the images.
build:
	scripts/build

# Push all the images.
push:
	scripts/push

# Pull latest containers.
pull:
	scripts/compose -f docker-compose.yml pull $(service)

# Start dependencies.
start:
	scripts/compose -f docker-compose.yml up -d --remove-orphans $(service)

# Stop dependencies.
stop:
	scripts/compose -f docker-compose.yml down --remove-orphans

# Logs from a service.
logs:
	scripts/compose -f docker-compose.yml logs -f $(service)

# Clean all unused docker images.
clean:
	scripts/clean image prune -a -f
