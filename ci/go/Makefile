build:
	docker build -t alexfalkowski/ci-golang:1.15 .

run:
	docker run -v /var/run/docker.sock:/var/run/docker.sock -v $(shell pwd)/docker-compose.yml:/etc/docker-compose.yml alexfalkowski/ci-golang:1.15

push: build
	docker push alexfalkowski/ci-golang:1.15
