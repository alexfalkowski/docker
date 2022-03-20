build:
	docker build -t alexfalkowski/ruby:3.0 .

push: build
	docker push alexfalkowski/ruby:3.0

lint:
	hadolint Dockerfile
