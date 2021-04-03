build:
	docker build -t alexfalkowski/go:1.16 .

push: build
	docker push alexfalkowski/go:1.16

lint:
	hadolint Dockerfile
