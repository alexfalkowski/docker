build:
	docker build -t alexfalkowski/kind:0.10 .

push: build
	docker push alexfalkowski/kind:0.10

lint:
	hadolint Dockerfile
