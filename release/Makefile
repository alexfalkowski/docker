build:
	docker build -t alexfalkowski/release:2.0 .

push: build
	docker push alexfalkowski/release:2.0

lint:
	hadolint Dockerfile
