build-docker:
	docker build -t alexfalkowski/diagram:1.0 .

push-docker: build-docker
	docker push alexfalkowski/diagram:1.0

lint-docker:
	hadolint Dockerfile
