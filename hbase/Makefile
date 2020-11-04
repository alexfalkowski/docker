build:
	docker build -t alexfalkowski/hbase:2.3.2 .

push: build
	docker push alexfalkowski/hbase:2.3.2

lint:
	docker run --rm -i hadolint/hadolint < Dockerfile
