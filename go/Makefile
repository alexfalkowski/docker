build:
	docker build -t alexfalkowski/go:1.21-3.2 .

push: build
	docker push alexfalkowski/go:1.21-3.2

lint:
	hadolint Dockerfile
