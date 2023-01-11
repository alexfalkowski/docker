build:
	docker build -t alexfalkowski/go:1.19-3.1.3 .

push: build
	docker push alexfalkowski/go:1.19-3.1.3

lint:
	hadolint Dockerfile
