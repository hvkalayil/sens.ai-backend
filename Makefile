.PHONY: all build run test clean lint migrateup migratedown migratenew

ifneq (,$(wildcard ./.env))
    include .env
    export
endif

CMD_PATH=cmd/server/main.go


sqlc:
	@echo "Generating sqlc code..."
	sqlc generate

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

migrateup:
	@echo "Running migration up..."
	migrate -path db/schema -database "$(DATABASE_URL)" -verbose up

migratedown:
	@echo "Running migration down..."
	migrate -path db/schema -database "$(DATABASE_URL)" -verbose down

migratenew:
	@echo "Creating new migration..."
	migrate create -ext sql -dir db/schema -seq $(name)
