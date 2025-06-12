# Build docker image.
build-docker:
	docker build -t alexfalkowski/$(IMAGE):$(VERSION) .

# Push built docker image.
push-docker:
	docker build -t alexfalkowski/$(IMAGE):$(VERSION) --push .

# Lint docker image.
lint-docker:
	hadolint Dockerfile
