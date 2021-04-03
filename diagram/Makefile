build:
	docker build -t alexfalkowski/diagram:1.0 .

push: build
	docker push alexfalkowski/diagram:1.0

lint:
	hadolint Dockerfile
