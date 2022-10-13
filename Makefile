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
	docker compose pull

# Start dependencies.
start:
	docker compose up -d --remove-orphans $(service)

# Stop dependencies.
stop:
	docker compose down --remove-orphans

# Logs from a service.
logs:
	docker compose logs -f $(service)

# Clean all unused docker images.
clean:
	docker image prune -a -f

# Verify the services.
verify:
	scripts/verify
