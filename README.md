# gotemplate

## Description

This is a template for a golang project with a postgres database. It includes a command-line tool (cli) and a grpc server for trying it.

## Architecture

It follows a hexagonal architecture. Nats for notifications.

## Requirements
you will need sudo for the commands
- go 1.24 https://go.dev/doc/install
`curl -L -o go1.24.1.linux-amd64.tar.gz https://go.dev/dl/go1.24.1.linux-amd64.tar.gz`
`rm -rf /usr/local/go && tar -C /usr/local -xzf go1.24.1.linux-amd64.tar.gz && rm go1.24.1.linux-amd64.tar.gz`
add to your ~/.bashrc `export PATH=$PATH:/usr/local/go/bin`
- golangci-lint https://golangci-lint.run/welcome/install/
`curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.64.6`
- docker
- docker-compose
- proto https://protobuf.dev/getting-started/gotutorial/ https://grpc.io/docs/protoc-installation/
`go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.30.0`
`go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0`
- pre-commit https://pre-commit.com/


## Dependencies

- [cobra](https://github.com/spf13/cobra) for the cli
- [testify](https://github.com/stretchr/testify) for testing

## Usage

This project comes with a makefile to ease the usage of it. It cover a wide range of commands:
- lint
- tests
- mocks
- build-docker
- run-docker-service
- run-docker-cli
- attach-docker

## Test

Creating user
```bash
curl -X POST http://localhost:8081/users \
    -H "Content-Type: application/json" \
    -d '{
        "first_name": "Alice",
        "last_name": "Bob",
        "nickname": "AB123",
        "email": "alice@bob.com",
        "country": "UK"
    }'
```


Retrieving user
```bash
curl -X GET http://localhost:8081/users/UUID
```