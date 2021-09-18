help: ## Display this help
	@ echo "Please use \`make <target>' where <target> is one of:"
	@ echo
	@ grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-10s\033[0m - %s\n", $$1, $$2}'
	@ echo

lint: ## Lint all the images
	scripts/lint

build: ## Build all the images
	scripts/build

push: ## Push all the images
	scripts/push

start: ## Start dependencies
	docker-compose up -d

stop: ## Stop dependencies
	docker-compose down
