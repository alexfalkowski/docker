build-docker:
	docker build -t alexfalkowski/go:1.22-ruby .

push-docker: build-docker
	docker push alexfalkowski/go:1.22-ruby

lint-docker:
	hadolint Dockerfile
