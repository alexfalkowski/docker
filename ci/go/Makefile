build:
	docker build -t go:1.15 .

run:
	docker run -v /var/run/docker.sock:/var/run/docker.sock -v $(shell pwd)/docker-compose.yml:/etc/docker-compose.yml go:1.15
