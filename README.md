# GoTemplate

## Description
GoTemplate is a Golang project with a PostgreSQL database. It includes a HTTP api for user operations. The project follows a clean architecture and utilizes NATS for notifications.

## Architecture
This project is designed following the **Clean Architecture** pattern, ensuring a decoupled, maintainable, and testable codebase. **Clean Architecture** organizes the application into distinct layers, each with explicit dependencies pointing inwards. This structure ensures that the core business logic remains independent of external systems, such as databases, frameworks, or user interfaces.

- **PostgreSQL**: Used as the primary database.
- **NATS**: Used for asynchronous messaging and notifications.

## Compatibility matrix

### app: 1.0.0

| Driver ↓ / Postgres → | 13 | 14 | 15 | 16 | 17 |
|:--------------------:|:---:|:---:|:---:|:---:|:---:|
| 1.10.9               | ✅  | ✅  | ✅  | ✅  | ✅  |

| Driver ↓ / NATS → | 2.3 | 2.4 | 2.5 | 2.6 | 2.7 | 2.8 | 2.9 | 2.10 |
|:--------------------:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|:---:|
| 1.23.5               | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | ✅  | 🟠  |

* 2.10 is a breaking change in nats https://docs.nats.io/release-notes/whats_new/whats_new_210

### Legend

| Symbol | Description |
|:------:|-------------|
| ✅     | Perfect match: all features are supported. Client and server versions have exactly the same features/APIs. |
| 🟠     | Forward compatibility: the client will work with the server, but not all new server features are supported. The server has features that the client library cannot use. |
| ❌     | Backward compatibility/Not applicable: the client has features that may not be present in the server. Common features will work, but some client APIs might not be available in the server. |
| -      | Not tested: this combination has not been verified or is not applicable. |

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

### 5. Install Pre-commit Hooks
Used for enforcing coding standards and automated checks, please refer to [Pre-commit](https://pre-commit.com/):
```sh
pip install pre-commit
pre-commit install
```

## Dependencies
This project includes the following key dependencies:
- [Testify](https://github.com/stretchr/testify) – Testing framework
- [Mux](https://github.com/gorilla/mux) – HTTP router
- [Viper](https://github.com/spf13/viper) – Configuration management
- [NATS](https://nats.io/) – Message broker

## Usage
GoTemplate provides a **Makefile** for easy management. Commonly used commands include:

### Build and run
```sh
make build   # Build Docker image for the app
make app     # Run the Docker container for the app
```

### Lint, test and generators
```sh
make lint   # Run Golangci-Lint, Goimports and Gofmt
make test   # Run unit and integration tests
make mocks  # Generate mock implementations
make sqlc   # Generate SQLC code
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
        "password": "password",
        "email": "alice@bob.com",
        "country": "UK"
    }'
```

### **2. Retrieve a User by ID**
```sh
curl -X GET http://localhost:8081/users/{user_id} \
    -H "Content-Type: application/json"
```

### **3. Remove a User**
```sh
curl -X DELETE http://localhost:8081/users/{user_id} \
    -H "Content-Type: application/json"
```

### **4. Get paginated Users list (with filters)**
```sh
curl -X GET "http://localhost:8081/users?country=USA&limit=10&offset=0 \
    -H "Content-Type: application/json"
```

### **5. Modify an existing User**
```sh
curl -X PATCH "http://localhost:8081/users/{user_id}" \
     -H "Content-Type: application/json" \
     -d '{"country": "USA"}'
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

### **Purge Services**
```sh
make purge
```
This stops and removes all running services (including containers, volumes, images and unused networks).

### **Accessing the container**
```sh
make attach
```

### Logs
To view logs, use the following command:
```sh
make logs
```

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
