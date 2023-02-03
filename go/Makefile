build:
	docker build -t alexfalkowski/go:1.20-3.1 .

push: build
	docker push alexfalkowski/go:1.20-3.1

lint:
	hadolint Dockerfile
