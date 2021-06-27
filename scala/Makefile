build:
	docker build -t alexfalkowski/scala:3.0 .

push: build
	docker push alexfalkowski/scala:3.0

lint:
	hadolint Dockerfile
