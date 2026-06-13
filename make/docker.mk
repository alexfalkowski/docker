DOCKER_IMAGE ?= alexfalkowski/$(IMAGE)

# Build docker image.
build-docker:
	@docker build -f Dockerfile -t $(DOCKER_IMAGE):$(VERSION) -t $(DOCKER_IMAGE) ..

# Push docker image.
push-docker:
	@docker push $(DOCKER_IMAGE):$(VERSION)
	@docker push $(DOCKER_IMAGE)

# Build and scan docker image.
test-docker: build-docker trivy-docker

# Scan docker image with Trivy.
trivy-docker:
	@trivy image --scanners vuln --pkg-types os --ignore-unfixed --db-repository public.ecr.aws/aquasecurity/trivy-db:2 --exit-code 1 --severity CRITICAL $(DOCKER_IMAGE):$(VERSION)

# Build, scan, and push docker image.
release-docker: test-docker push-docker

# Build platform docker image.
build-platform-docker:
	@docker build --platform linux/$(platform) -f Dockerfile -t $(DOCKER_IMAGE):$(VERSION).$(platform) ..

# Push platform docker image.
push-platform-docker:
	@docker push $(DOCKER_IMAGE):$(VERSION).$(platform)

# Build and scan platform docker image.
test-platform-docker: build-platform-docker trivy-platform-docker

# Scan platform docker image with Trivy.
trivy-platform-docker:
	@trivy image --scanners vuln --pkg-types os --ignore-unfixed --db-repository public.ecr.aws/aquasecurity/trivy-db:2 --exit-code 1 --severity CRITICAL $(DOCKER_IMAGE):$(VERSION).$(platform)

# Build, scan, and push platform docker image.
release-platform-docker: test-platform-docker push-platform-docker

manifest-platform-version-docker:
	@docker manifest create $(DOCKER_IMAGE):$(VERSION) --amend $(DOCKER_IMAGE):$(VERSION).amd64 --amend $(DOCKER_IMAGE):$(VERSION).arm64
	@docker manifest push $(DOCKER_IMAGE):$(VERSION)

manifest-platform-latest-docker:
	@docker manifest create $(DOCKER_IMAGE) --amend $(DOCKER_IMAGE):$(VERSION).amd64 --amend $(DOCKER_IMAGE):$(VERSION).arm64
	@docker manifest push $(DOCKER_IMAGE)

# Create a platform manifest.
manifest-platform-docker: manifest-platform-version-docker manifest-platform-latest-docker

# Lint docker image.
lint-docker:
	@hadolint Dockerfile
