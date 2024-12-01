# Dealls Dating App

## Overview

This service provides functionality for deals Dating App. The service is built using the Go programming language and follows a layered architecture pattern. The service is designed to be scalable, maintainable, and testable.

---

## Project Structure

The project is organized into the following directories:

```
.
├── api               # API entry points for the service
│   └── api.go        # Routes and API definitions
├── config            # Configuration files and settings
│   └── app.toml.dist # Configuration template
├── handler           # HTTP handlers for the service
├── infra             # Infrastructure-related code
│   └── infra.go
├── interfaces        # Interface definitions and mocks
│   └── mocks         # Mock implementations for testing
├── manager           # Managers for repository and service layers
├── middleware        # Middleware for request validation and token management
├── model             # Data models and constants
├── repo              # Repository implementations for data access
├── resources         # DTOs (Data Transfer Objects)
│   ├── request       # Request payload definitions
│   └── response      # Response payload definitions
├── schema            # Database schema and migrations
├── service           # Business logic for the application
├── unittest          # Unit tests for services and utilities
├── utils             # Utility functions and helpers
├── go.mod            # Go module definition
├── go.sum            # Dependencies lockfile
└── .gitignore        # Files and directories to be ignored by Git
```

---

## Requirements

Ensure the following are installed on your machine:

- Go (>= 1.20)
- Git
- Docker (for containerized deployment)
- GolangCI-Lint (for linting and code quality checks)

---

## Running the Service

1. Clone the repository:
   ```bash
   git clone https://github.com/dwiangraenih/dealls.git
   cd dealls
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   go mod vendor
   ```

3. Configure the application:
    - Copy `app.toml.dist` to `app.toml`.
    - Update `app.toml` with your database and secret configurations.

4. Prepare the database:
    - Copy and execute the schema file (`schema/0001_init.up.sql`) in your database.

5. Run the project:
   ```bash
   go run main.go api
   ```

---

## Test Service
Use the following command to test the service:

```bash
go test $(go list ./... | grep -v '/mocks$') -coverprofile=coverage.out
```

---

## Update Mocks
Use the following command to update mocks:

```bash
mockgen -source=your-source-code.go -destination=your-destination-code.go
```

---

## Docker Setup

### Building the Docker Image
Build the Docker image for the service:
```bash
docker build -t dealls_dating_app .
```

### Running the Docker Container
Run the service in a container:
```bash
docker run -p 8010:8010 dealls_dating_app
```

### Using Docker Compose
If you have a `docker-compose.yml` file, use the following command to run the application along with its dependencies:
```bash
docker-compose up --build
```

---

## GolangCI-Lint

### Installing GolangCI-Lint
Install GolangCI-Lint using Homebrew:
```bash
brew install golangci-lint
```

### Running GolangCI-Lint
To lint the project and enforce coding standards:
```bash
golangci-lint run
```

### Fixing Issues
To automatically fix issues detected by the linter:
```bash
golangci-lint run --fix
```

## Notes

- Ensure your Go version matches or exceeds the minimum required version (1.20).
- Logs will provide detailed error messages in case of runtime issues.
- Docker and GolangCI-Lint are recommended for maintaining a consistent development and deployment workflow.

For further assistance, please contact the repository maintainer.
