.PHONY: all build run test clean lint

CMD_PATH=cmd/server/main.go


build:
	@echo "Building..."
	go build -o bin/ $(CMD_PATH)

run:
	make docs
	@echo "Running..."
	go run $(CMD_PATH)

test:
	@echo "Testing..."
	go test ./... -v
check:
	@echo "Checking..."
	pre-commit run --all-files

deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

docs:
	@echo "Generating Swagger docs..."
	swag init -g cmd/server/main.go
