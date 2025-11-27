.PHONY: all build run test clean lint

APP_NAME=sens-ai-backend
CMD_PATH=cmd/server/main.go

all: build

build:
	@echo "Building..."
	go build -o bin/$(APP_NAME) $(CMD_PATH)

run:
	@echo "Running..."
	go run $(CMD_PATH)

test:
	@echo "Testing..."
	go test ./... -v

clean:
	@echo "Cleaning..."
	rm -rf bin/

lint:
	@echo "Linting..."
	golangci-lint run

deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy
