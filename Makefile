help: ## Display this help
	@ echo "Please use \`make <target>' where <target> is one of:"
	@ echo
	@ grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-10s\033[0m - %s\n", $$1, $$2}'
	@ echo

lint: ## Lint all the images
	make -C go lint
	make -C hbase lint
	make -C java lint
	make -C kotlin lint
	make -C release lint
	make -C ruby lint
	make -C scala lint
	make -C diagram lint

build: ## Build all the images
	make -C go build
	make -C hbase build
	make -C java build
	make -C kotlin build
	make -C release build
	make -C ruby build
	make -C scala build
	make -C diagram build

push: ## Push all the images
	make -C go push
	make -C hbase push
	make -C java push
	make -C kotlin push
	make -C release push
	make -C ruby push
	make -C scala push
	make -C diagram push
