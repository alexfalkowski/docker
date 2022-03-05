build:
	docker build -t alexfalkowski/java:17 .

push: build
	docker push alexfalkowski/java:17
lint:
	hadolint Dockerfile
