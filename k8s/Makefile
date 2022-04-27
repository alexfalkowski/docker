build:
	docker build -t alexfalkowski/k8s:1.0 .

push: build
	docker push alexfalkowski/k8s:1.0

lint:
	hadolint Dockerfile
