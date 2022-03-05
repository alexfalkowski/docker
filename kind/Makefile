build:
	docker build -t alexfalkowski/kind:0.11 .

push: build
	docker push alexfalkowski/kind:0.11

lint:
	hadolint Dockerfile
