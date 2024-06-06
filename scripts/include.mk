# Build docker image.
build-docker:
	docker build -t alexfalkowski/$(IMAGE):$(VERSION) .
	slim build --preserve-path-file ../scripts/preserved_paths.txt --http-probe-off --continue-after=1 --target alexfalkowski/$(IMAGE):$(VERSION) --tag alexfalkowski/$(IMAGE):$(VERSION)-slim

# Push built docker image.
push-docker: build-docker
	docker push alexfalkowski/$(IMAGE) --all-tags

# Lint docker image.
lint-docker:
	hadolint Dockerfile
