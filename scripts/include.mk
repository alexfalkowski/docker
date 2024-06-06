# Build docker image.
build-docker:
	docker build --compress -t alexfalkowski/$(IMAGE) .
	slim build --http-probe-off --continue-after=1 --target alexfalkowski/$(IMAGE) --tag alexfalkowski/$(IMAGE)-slim

# Push built docker image.
push-docker: build-docker
	docker push alexfalkowski/$(IMAGE) alexfalkowski/$(IMAGE)-slim

# Lint docker image.
lint-docker:
	hadolint Dockerfile
