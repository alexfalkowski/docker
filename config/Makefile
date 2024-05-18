build-docker:
	docker build -t alexfalkowski/config:1.0 .

push-docker: build-docker
	docker push alexfalkowski/config:1.0

lint-docker:
	hadolint Dockerfile
