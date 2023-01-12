build:
	docker build -t alexfalkowski/release:3.1 .

push: build
	docker push alexfalkowski/release:3.1

lint:
	hadolint Dockerfile
