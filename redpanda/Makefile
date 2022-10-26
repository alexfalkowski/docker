build:
	docker build -t alexfalkowski/redpanda:1.0 .

push: build
	docker push alexfalkowski/redpanda:1.0

lint:
	hadolint Dockerfile
