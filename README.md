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

---

## Running the Service

- Clone [repository](https://github.com/dwiangraenih/dealls.git)
- Install Dependecies `go mod tidy` and `go mod vendor`
- Copy file `app.toml.dist` to be `app.toml`
- Check pdf file for hidden key
- Copy and run the schema file to your database
- Run Project with `go run main.go api`

## Test Service
Use the following command to test the service:

```bash
go test $(go list ./... | grep -v '/mocks$') -coverprofile=coverage.out
```

## Update Mocks
Use the following command to update mocks:

```bash
mockgen -source=your-source-code.go -destination=your-destination-code.go
```

---

## Notes

- Ensure your Go version matches or exceeds the minimum required version (1.20).
- Logs will provide detailed error messages in case of runtime issues.

For further assistance, please contact the repository maintainer.