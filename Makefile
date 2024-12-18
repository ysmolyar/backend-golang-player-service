.PHONY: build run test clean fmt lint help setup

# Default target
.DEFAULT_GOAL := help

# Go related variables
BINARY_NAME=player-service
MAIN_FILE=cmd/server/main.go

# Build the application
build: ## Build the application
	go build -o bin/$(BINARY_NAME) $(MAIN_FILE)

# Run the application
run: ## Run the application
	./bin/$(BINARY_NAME)

# Clean build files
clean: ## Clean build files
	rm -rf bin/
	go clean

# Run tests
test: ## Run tests
	go test -v ./...

# Format code
fmt: ## Format code using gofmt
	go fmt ./...

# Run linter
lint: ## Run golangci-lint
	golangci-lint run

# Install dependencies
deps: ## Install dependencies
	go mod download
	go mod tidy

# Show help
help: ## Show this help message
	@echo 'Usage:'
	@echo '  make <target>'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*##"; printf "\033[36m"} /^[a-zA-Z_-]+:.*?##/ { printf "  %-15s %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST) 