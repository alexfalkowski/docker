build:
	docker build -t alexfalkowski/java:15 .

push: build
	docker push alexfalkowski/java:15
lint:
	hadolint Dockerfile
