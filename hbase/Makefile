build:
	docker build -t alexfalkowski/hbase:2.4.2 .

push: build
	docker push alexfalkowski/hbase:2.4.2

lint:
	hadolint Dockerfile
