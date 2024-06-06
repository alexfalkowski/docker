# Build docker image.
build-docker:
	docker build -t alexfalkowski/$(IMAGE):$(VERSION) .

# Push built docker image.
push-docker: build-docker
	docker push alexfalkowski/$(IMAGE) --all-tags

# Lint docker image.
lint-docker:
	hadolint Dockerfile
