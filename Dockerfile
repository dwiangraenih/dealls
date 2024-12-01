# Use official Golang image for build
FROM golang:1.21 AS builder

# Set working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum for dependency installation
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application binary
RUN go build -o main .

# Use minimal base image for running the app
FROM debian:bullseye-slim

# Set working directory for runtime
WORKDIR /app

# Copy the compiled binary from the builder
COPY --from=builder /app/main .

# Expose the application port
EXPOSE 8010

# Command to run the app
CMD ["./main", "api"]
