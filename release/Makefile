build-docker:
	docker build -t alexfalkowski/release:3.1 .

push-docker: build-docker
	docker push alexfalkowski/release:3.1

lint-docker:
	hadolint Dockerfile
