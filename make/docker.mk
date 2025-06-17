# Build docker image.
build-docker:
	docker build -t alexfalkowski/$(IMAGE):$(VERSION) .

# Push built docker image.
push-docker:
	docker build -t alexfalkowski/$(IMAGE):$(VERSION) -t alexfalkowski/$(IMAGE) --push .

# Build platform docker image.
build-platform-docker:
	docker build -t alexfalkowski/$(IMAGE):$(VERSION).$(platform) .

# Push built platform docker image.
push-platform-docker:
	docker build -t alexfalkowski/$(IMAGE):$(VERSION).$(platform) --push .

manifest-platform-version-docker:
	docker manifest create alexfalkowski/$(IMAGE):$(VERSION) --amend alexfalkowski/$(IMAGE):$(VERSION).amd64 --amend alexfalkowski/$(IMAGE):$(VERSION).arm64
	docker manifest push alexfalkowski/$(IMAGE):$(VERSION)

manifest-platform-latest-docker:
	docker manifest create alexfalkowski/$(IMAGE) --amend alexfalkowski/$(IMAGE):$(VERSION).amd64 --amend alexfalkowski/$(IMAGE):$(VERSION).arm64
	docker manifest push alexfalkowski/$(IMAGE)

# Create a platform manifest.
manifest-platform-docker: manifest-platform-version-docker manifest-platform-latest-docker

# Lint docker image.
lint-docker:
	hadolint Dockerfile
