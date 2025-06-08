# Makefile for EvoGarden development tasks

GOFMT=gofmt
GOVET=go vet
GOTEST=go test

.PHONY: all build run fmt lint test clean

# Default target
all: build

# Build the evogarden executable
build: fmt lint
	go build -o bin/evogarden ./cmd/evogarden

# Run the application
run:
	go run ./cmd/evogarden

# Format Go source files
fmt:
	$(GOFMT) -w $(shell git ls-files '*.go')

# Lint the code
lint:
	$(GOVET) ./...

# Execute unit tests
test:
	$(GOTEST) ./...

# Clean build artifacts
clean:
	rm -rf bin
