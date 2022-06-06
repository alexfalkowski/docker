build:
	docker build -t alexfalkowski/clang:1.0 .

push: build
	docker push alexfalkowski/clang:1.0

lint:
	hadolint Dockerfile
