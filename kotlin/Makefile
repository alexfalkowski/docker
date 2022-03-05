build:
	docker build -t alexfalkowski/kotlin:1.6 .

push: build
	docker push alexfalkowski/kotlin:1.6

lint:
	hadolint Dockerfile
