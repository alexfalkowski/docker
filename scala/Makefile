build:
	docker build -t alexfalkowski/scala:2.13 .

push: build
	docker push alexfalkowski/scala:2.13

lint:
	hadolint Dockerfile
