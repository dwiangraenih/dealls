# Variables
APP_NAME = dealls_dating_app
DOCKER_IMAGE = $(APP_NAME)
BINARY_NAME = main

# Build binary
build:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME) .

# Run tests
test:
	go test ./... -v

# Build Docker image
docker-build:
	docker build -t $(DOCKER_IMAGE) .

# Run Docker container
docker-run:
	docker run --rm -p 8080:8080 --name $(APP_NAME) $(DOCKER_IMAGE)

# Clean up
clean:
	rm -f $(BINARY_NAME)

# Stop Docker container
docker-stop:
	docker stop $(APP_NAME) || true

# Remove Docker image
docker-clean:
	docker rmi $(DOCKER_IMAGE) || true

# Full rebuild (clean, build binary, build Docker image)
rebuild: clean build docker-build
