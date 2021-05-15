build:
	docker build -t alexfalkowski/kotlin:1.5 .

push: build
	docker push alexfalkowski/kotlin:1.5

lint:
	hadolint Dockerfile
