build:
	docker build -t alexfalkowski/terraform:1.0 .

push: build
	docker push alexfalkowski/terraform:1.0

lint:
	hadolint Dockerfile
