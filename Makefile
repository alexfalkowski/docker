# Lint all the images
lint:
	scripts/lint

# Build all the images
build:
	scripts/build

# Push all the images
push:
	scripts/push

# Start dependencies
start:
	docker-compose up -d

# Stop dependencies
stop:
	docker-compose down
