# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install required system dependencies
RUN apk add --no-cache git

# Copy go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server/main.go

# Production stage
FROM alpine:latest

WORKDIR /root/

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose the application port
EXPOSE 3000

# Command to run the executable
CMD ["./main"]
