# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install required system dependencies
RUN apk add --no-cache git

# Copy go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest

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
RUN apk --no-cache add ca-certificates

# Copy the binary from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs

# Expose the application port
EXPOSE 3000

# Command to run the executable
CMD ["./main"]
