build:
	docker build -t alexfalkowski/release:1.0 .

push: build
	docker push alexfalkowski/release:1.0

lint:
	docker run --rm -i hadolint/hadolint < Dockerfile
