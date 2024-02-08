build:
	docker build -t alexfalkowski/go:1.22-ruby .

push: build
	docker push alexfalkowski/go:1.22-ruby

lint:
	hadolint Dockerfile
