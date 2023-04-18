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
	docker compose -f $(kind)-docker-compose.yml pull $(service)

# Start dependencies.
start:
	docker compose -f $(kind)-docker-compose.yml up -d --remove-orphans $(service)

# Stop dependencies.
stop:
	docker compose -f $(kind)-docker-compose.yml down --remove-orphans

# Logs from a service.
logs:
	docker compose -f $(kind)-docker-compose.yml logs -f $(service)

# Clean all unused docker images.
clean:
	docker image prune -a -f

# Verify the services.
verify:
	scripts/verify
