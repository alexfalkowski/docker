build:
	docker build -t alexfalkowski/go:1.17 .

push: build
	docker push alexfalkowski/go:1.17

lint:
	hadolint Dockerfile
