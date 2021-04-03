build:
	docker build -t alexfalkowski/java:16 .

push: build
	docker push alexfalkowski/java:16
lint:
	hadolint Dockerfile
