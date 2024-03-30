build-docker:
	docker build -t alexfalkowski/ruby:3.3 .

push-docker: build-docker
	docker push alexfalkowski/ruby:3.3

lint-docker:
	hadolint Dockerfile
