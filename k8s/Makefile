build-docker:
	docker build -t alexfalkowski/k8s:1.0 .

push-docker: build-docker
	docker push alexfalkowski/k8s:1.0

lint-docker:
	hadolint Dockerfile
