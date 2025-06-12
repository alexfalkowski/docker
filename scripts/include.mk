multi-builder-docker:
	docker buildx create --name docker-multi-builder --driver docker-container

# Build docker image.
build-docker: multi-builder-docker
	docker buildx build --builder docker-multi-builder --platform $(platform) -t alexfalkowski/$(IMAGE):$(VERSION) .

# Push built docker image.
push-docker: multi-builder-docker
	docker buildx build --builder docker-multi-builder --platform $(platform) -t alexfalkowski/$(IMAGE):$(VERSION) --push .

# Lint docker image.
lint-docker:
	hadolint Dockerfile
