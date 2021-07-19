build:
	docker build -t alexfalkowski/hbase:2.4.4 .

push: build
	docker push alexfalkowski/hbase:2.4.4

lint:
	hadolint Dockerfile
