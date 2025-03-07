.DEFAULT_GOAL := help

.PHONY: help
help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: lint
lint: ## Run golangci-lint, goimports and gofmt
	golangci-lint run --config .golangci.yml ./... && go tool goimports -w . && gofmt -s -w -e -d .

.PHONY: tests
tests: ## Executes api tests
	go test ./... --tags=integration,unit -coverpkg=./...

.PHONY: mocks
mocks: ## Generates mocks
	go generate ./...

.PHONY: proto
proto: ## Generates proto files
	protoc -I=./proto --go_out=proto --go_opt=paths=source_relative --go-grpc_out=./proto --go-grpc_opt=paths=source_relative proto/user.proto

.PHONY: build
build: ## Build the Docker image
	docker build -t app -f Dockerfile .

.PHONY: server
server: ## Build and Run the Docker server
	docker-compose up -d app

.PHONY: logs
logs: ## Show the Docker logs
	docker-compose logs -f

.PHONY: local-cli
local-cli: ## Build the local CLI binary
ifeq ($(OS),Windows_NT)
	go build -o cli.exe ./cmd/cli
else
	go build -o cli ./cmd/cli
endif

.PHONY: cli
cli: ## Run the Docker CLI
	docker up -it --rm app /app/cli $(ARGS)

.PHONY: attach
attach: ## Attach to the Docker container
	docker-compose run -it --rm --entrypoint /bin/sh app

.PHONY: purge
purge: ## Purge all Docker containers and images
	docker rm -f `docker ps -a -q` || true
	docker rmi -f `docker images -q` || true
	docker volume prune -f
	docker network prune -f