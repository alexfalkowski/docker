build:
	docker build -t alexfalkowski/go:1.19 .

push: build
	docker push alexfalkowski/go:1.19

lint:
	hadolint Dockerfile
