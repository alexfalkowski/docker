build:
	docker build -t alexfalkowski/kotlin:1.4 .

push: build
	docker push alexfalkowski/kotlin:1.4

lint:
	hadolint Dockerfile
