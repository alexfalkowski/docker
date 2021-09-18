build:
	docker build -t alexfalkowski/hbase:2.4.6 .

push: build
	docker push alexfalkowski/hbase:2.4.6

lint:
	hadolint Dockerfile
