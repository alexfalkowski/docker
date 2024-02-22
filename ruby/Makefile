build:
	docker build -t alexfalkowski/ruby:3.3 .

push: build
	docker push alexfalkowski/ruby:3.3

lint:
	hadolint Dockerfile
