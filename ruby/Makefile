build:
	docker build -t alexfalkowski/ruby:2.7 .

push: build
	docker push alexfalkowski/ruby:2.7

lint:
	docker run --rm -i hadolint/hadolint < Dockerfile
