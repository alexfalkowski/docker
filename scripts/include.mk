# Build docker image.
build-docker:
	docker build --compress -t alexfalkowski/$(IMAGE) .
	slim build --sensor-ipc-mode proxy --sensor-ipc-endpoint 172.17.0.1 --http-probe=false alexfalkowski/$(IMAGE)

# Push built docker image.
push-docker: build-docker
	docker push alexfalkowski/$(IMAGE)

# Lint docker image.
lint-docker:
	hadolint Dockerfile
