build:
	docker build -t alexfalkowski/release:3.0 .

push: build
	docker push alexfalkowski/release:3.0

lint:
	hadolint Dockerfile
