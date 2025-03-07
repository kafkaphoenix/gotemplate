# GoTemplate

## Description
GoTemplate is a Golang project template with a PostgreSQL database. It includes a command-line tool (CLI) and a gRPC server for managing user operations. The project follows a hexagonal architecture and utilizes NATS for notifications.

## Architecture
This project is designed following the **Clean Architecture** pattern, ensuring a decoupled, maintainable, and testable codebase. Clean Architecture organizes the application into distinct layers, each with explicit dependencies pointing inwards. This structure ensures that the core business logic remains independent of external systems, such as databases, frameworks, or user interfaces.

- **PostgreSQL**: Used as the primary database.
- **NATS**: Used for asynchronous messaging and notifications.
- **gRPC**: Supports efficient inter-service communication.
- **Cobra**: CLI framework for command-line interactions.
- **Testify**: Used for unit testing.

## Requirements
To run this project, please follow the steps listed below:

### 1. Install Go 1.24
Follow the official [Go installation guide](https://go.dev/doc/install) or use the following commands:

```sh
curl -L -o go1.24.1.linux-amd64.tar.gz https://go.dev/dl/go1.24.1.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.24.1.linux-amd64.tar.gz && rm go1.24.1.linux-amd64.tar.gz
```
Add the Go binary path to your shell profile (`~/.bashrc` or `~/.zshrc`):
```sh
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin
```

### 2. Install GolangCI-Lint
A linter for Go projects, more info please refer to [GolangCI-Lint](https://golangci-lint.run/welcome/install/):
```sh
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.64.6
```

### 3. Install Docker & Docker Compose
Ensure that Docker and Docker Compose are installed:
- [Docker Installation](https://docs.docker.com/get-docker/)
- [Docker Compose Installation](https://docs.docker.com/compose/install/)

### 4. Install Protocol Buffers (proto) and gRPC
Install Protocol Buffers for gRPC communication:
```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.30.0
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
```
For additional setup, refer to:
- [Protobuf Setup](https://protobuf.dev/getting-started/gotutorial/)
- [gRPC Setup](https://grpc.io/docs/protoc-installation/)

### 5. Install Pre-commit Hooks
Used for enforcing coding standards and automated checks, please refer to [Pre-commit](https://pre-commit.com/):
```sh
pip install pre-commit
pre-commit install
```

## Dependencies
This project includes the following key dependencies:
- [Cobra](https://github.com/spf13/cobra) – CLI framework
- [Testify](https://github.com/stretchr/testify) – Testing framework
- [Mux](https://github.com/gorilla/mux) – HTTP router
- [Viper](https://github.com/spf13/viper) – Configuration management
- [GORM](https://gorm.io/) – ORM library
- [NATS](https://nats.io/) – Message broker

## Usage
GoTemplate provides a **Makefile** for easy management. Commonly used commands include:

### Build and Run
```sh
make build   # Build Docker image
make server  # Run the backend service in Docker
make cli     # Run CLI inside the container
```

### Lint, Test and generators
```sh
make lint   # Run Golangci-Lint, Goimports and Gofmt
make test   # Run unit and integration tests
make mocks  # Generate mock implementations
make proto  # Generate gRPC code
```

## API Endpoints
### **1. Create a User**
```sh
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

### **2. Retrieve a User by ID**
```sh
curl -X GET http://localhost:8081/users/{user_id} \
    -H "Content-Type: application/json"
```

### **3. Delete a User**
```sh
curl -X DELETE http://localhost:8081/users/{user_id} \
    -H "Content-Type: application/json"
```

### **4. Get Users List (with filters)**
```sh
curl -X GET "http://localhost:8081/users?country=USA&limit=10&offset=0 \
    -H "Content-Type: application/json"
```

### **5. Update User Information**
```sh
curl -X PUT "http://localhost:8081/users/{user_id}" \
     -H "Content-Type: application/json" \
     -d '{"first_name": "UpdatedName", "last_name": "UpdatedLast", "email": "updated.email@example.com", "country": "Canada"}'
```

## Running in Docker
### **Start Services**
```sh
make server
```
This starts the following services:
- PostgreSQL (database)
- NATS (message broker)
- GoTemplate (application server)

### **Stopping Services**
```sh
make purge
```
This stops and removes all running services (includind containers, volumes, images and networks).

### **Accessing the app container**
```sh
make attach
```

### Logging
To view logs, use the following command:
```sh
make logs
```

### Command Line Interface (CLI)
The CLI is used for managing user operations. To run the CLI, use the following command:
```sh
make local-cli
```

#### CLI Commands
- **Create User**: `./cli user create --first-name "Alice" --last-name "Bob" --nickname "AB123" --email "alice@bob.com" --country "UK"`
- **Get User**: `./cli user get --id {user_id}`

## Contributing
To contribute:
1. Fork the repository.
2. Clone your fork.
3. Create a new branch (`git checkout -b feature-branch`).
4. Commit changes (`git commit -m "Add new feature"`).
5. Push to the branch (`git push origin feature-branch`).
6. Open a pull request.

## License
This project is licensed under the MIT License. See `LICENSE` for details.

