# Build docker image.
build-docker:
	docker build -t alexfalkowski/$(IMAGE)-$(platform):$(VERSION) .

# Push built docker image.
push-docker:
	docker build -t alexfalkowski/$(IMAGE)-$(platform):$(VERSION) --push .

# Create a manifest.
manifest-docker:
	docker manifest create alexfalkowski/$(IMAGE)-$(platform):$(VERSION) --amend alexfalkowski/$(IMAGE)-amd64:$(VERSION) --amend alexfalkowski/$(IMAGE)-amd64:$(VERSION)
	docker manifest push alexfalkowski/$(IMAGE)-$(platform):$(VERSION)

# Lint docker image.
lint-docker:
	hadolint Dockerfile
