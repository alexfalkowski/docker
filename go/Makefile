build:
	docker build -t alexfalkowski/go:1.19-3.2.0 .

push: build
	docker push alexfalkowski/go:1.19-3.2.0

lint:
	hadolint Dockerfile
