.DEFAULT_GOAL := help

.PHONY: help
help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: lint
lint: ## Run golangci-lint, goimports and gofmt
	golangci-lint run --config .golangci.yml ./... && go tool goimports -w . && gofmt -s -w -e -d .

.PHONY: tests
tests: ## Executes tests
	go test ./... --tags=integration,unit -coverpkg=./...

.PHONY: mocks
mocks: ## Generates mocks
	go generate ./...

.PHONY: sqlc
sqlc: ## Generate SQLC code
	go tool sqlc generate

.PHONY: build
build: ## Build the Docker image
	docker build -t app -f Dockerfile .

.PHONY: server
server: ## Build and Run the Docker server
	docker-compose up -d app

.PHONY: logs
logs: ## Show all the Docker logs
	docker-compose logs -f

.PHONY: attach
attach: ## Attach to the Docker container
ifeq ($(OS),Windows_NT)
	docker-compose run -it --rm --entrypoint sh app
else
	docker-compose run -it --rm --entrypoint /bin/sh app
endif

.PHONY: purge
purge: ## Purge Docker containers, images, volumes and unused networks
	-docker rm -f `docker ps -a -q`
	-docker rmi -f `docker images -q`
	-docker volume rm `docker volume ls -q`
	docker network prune -f
