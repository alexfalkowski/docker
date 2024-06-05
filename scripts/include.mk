# Build docker image.
build-docker:
	docker build --compress -t alexfalkowski/$(IMAGE) .

# Push built docker image.
push-docker: build-docker
	docker push alexfalkowski/$(IMAGE)

# Lint docker image.
lint-docker:
	hadolint Dockerfile
