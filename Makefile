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
	docker-compose -f docker-compose.yml pull $(service)

# Start dependencies.
start:
	docker-compose -f docker-compose.yml up -d --remove-orphans $(service)

# Stop dependencies.
stop:
	docker-compose -f docker-compose.yml down --remove-orphans

# Logs from a service.
logs:
	docker-compose -f docker-compose.yml logs -f $(service)

# Clean all unused docker images.
clean:
	docker image prune -a -f

# Start procs.
start-procs:
	overmind start -D

# Start procs.
stop-procs:
	overmind quit

# Logs from a proc.
logs-procs:
	overmind echo

# Complile service.
compile:
	./scripts/compile $(service)

# Create certificates.
create-certs:
	mkcert -key-file config/certs/key.pem -cert-file config/certs/cert.pem localhost host.containers.internal
	mkcert -client -key-file config/certs/client-key.pem -cert-file config/certs/client-cert.pem localhost host.containers.internal

# Recompile all the proces.
procs:
	scripts/procs

# Verify the services.
verify:
	scripts/verify

# Verify the auth.
auth:
	scripts/auth

# Load test services.
load:
	scripts/load
