build:
	docker build -t alexfalkowski/ruby:3.1 .

push: build
	docker push alexfalkowski/ruby:3.1

lint:
	hadolint Dockerfile
