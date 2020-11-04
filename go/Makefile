build:
	docker build -t alexfalkowski/go:1.15 .

push: build
	docker push alexfalkowski/go:1.15

lint:
	docker run --rm -i hadolint/hadolint < Dockerfile
