# gotemplate

## Description

This is a template for a golang project with a postgres database. It includes a command-line tool (cli) and a grpc server for trying it.

## Architecture

It follows a hexagonal architecture. Nats for notifications.

## Requirements

- go 1.24
- golangci-lint https://golangci-lint.run/welcome/install/
- docker
- docker-compose

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


