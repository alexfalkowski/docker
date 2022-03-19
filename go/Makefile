build:
	docker build -t alexfalkowski/go:1.18 .

push: build
	docker push alexfalkowski/go:1.18

lint:
	hadolint Dockerfile
