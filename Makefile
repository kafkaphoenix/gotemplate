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

.PHONY: build-docker
build-docker: ## Build the Docker image
	docker build -t app -f Dockerfile .

.PHONY: run-docker-server
run-docker-server: ## Run the Docker server
	docker run -it --rm app /app/server

.PHONY: run-docker-cli
run-docker-cli: ## Run the Docker CLI
	docker run -it --rm app /app/cli $(ARGS)