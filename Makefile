# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Binary names
BINARY_NAME=./bin/api

all: test build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/api/main.go
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./cmd/api/main.go
	./$(BINARY_NAME)
