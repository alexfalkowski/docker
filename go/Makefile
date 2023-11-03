build:
	docker build -t alexfalkowski/go:1.21-ruby .

push: build
	docker push alexfalkowski/go:1.21-ruby

lint:
	hadolint Dockerfile
