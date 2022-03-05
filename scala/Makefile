build:
	docker build -t alexfalkowski/scala:3.1 .

push: build
	docker push alexfalkowski/scala:3.1

lint:
	hadolint Dockerfile
