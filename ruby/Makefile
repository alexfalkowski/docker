build:
	docker build -t alexfalkowski/ruby:3.2 .

push: build
	docker push alexfalkowski/ruby:3.2

lint:
	hadolint Dockerfile
