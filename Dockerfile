# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install required system dependencies
RUN apk add --no-cache git gcc musl-dev

# Copy go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Install golang-migrate CLI
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Copy the source code
COPY . .

# Generate Swagger docs
RUN swag init -g ./cmd/server/main.go --output ./docs

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server/main.go

# Production stage
FROM alpine:latest AS runner

WORKDIR /root/

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates bash

# Copy the binary from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs

# Copy golang-migrate binary
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate

# Copy migration files so they can be run in the container
COPY --from=builder /app/db/schema ./db/schema

# Expose the application port
EXPOSE 3000

# Command to run migrations and then start the executable
CMD ["sh", "-c", "migrate -path db/schema -database ${DATABASE_URL} up && ./main"]
