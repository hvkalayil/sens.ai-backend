# Sens.ai Backend

Go REST API project using Fiber.

## Prerequisites

- Go 1.20 or higher
- Python 3.7+ (for pre-commit)
- Git

## Setup

### 1. Clone the Repository
```bash
git clone https://github.com/hvkalayil/sens.ai-backend.git
cd sens.ai-backend
```

### 2. Install Go Dependencies
```bash
go mod download
```

Or using the Makefile:
```bash
make deps
```

### 3. Set Up Environment Variables
Create a `.env` file in the root directory:
```bash
cp .env.example .env
```

Edit `.env` with your configuration:
```env
HOST=localhost
PORT=3000
LOG_LEVEL=debug
```

### 4. Install Required Go Tools

Install the following tools for code quality and documentation:

```bash
# Install goimports (for import formatting)
go install golang.org/x/tools/cmd/goimports@latest

# Install golangci-lint (for linting)
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Install swag (for API documentation)
go install github.com/swaggo/swag/cmd/swag@latest
```

**Important for Windows users:** Add Go's bin directory to your PATH:
```powershell
# Add to your PowerShell profile or set permanently in System Environment Variables
$env:PATH += ";$((go env GOPATH))\bin"
```

Or permanently add `C:\Users\<YourUsername>\go\bin` to your system PATH.

### 5. Install Pre-commit

Install pre-commit using pip:
```bash
pip install pre-commit
```

Or on Windows with Python:
```powershell
python -m pip install pre-commit
```

### 6. Install Pre-commit Hooks
```bash
pre-commit install
```

This will automatically run code formatting, linting, and other checks before each commit.

### 7. Generate API Documentation
```bash
swag init -g cmd/server/main.go
```

Or using the Makefile:
```bash
make docs
```

## Usage

### Run Locally
```bash
make run
```

### Build
```bash
make build
```

### Test
```bash
make test
```

### Lint
```bash
make lint
```

## Pre-commit Hooks

This project uses pre-commit hooks to ensure code quality. The following checks run automatically before each commit:

- **trailing-whitespace**: Removes trailing whitespace
- **end-of-file-fixer**: Ensures files end with a newline
- **check-yaml**: Validates YAML files
- **check-added-large-files**: Prevents committing large files
- **go fmt**: Formats Go code
- **go vet**: Examines Go source code and reports suspicious constructs
- **goimports**: Formats imports and organizes them
- **golangci-lint**: Runs comprehensive Go linting
- **swag init**: Regenerates API documentation

### Running Pre-commit Manually

To run all hooks on all files:
```bash
pre-commit run --all-files
```

To run a specific hook:
```bash
pre-commit run golangci-lint --all-files
```

### Troubleshooting

**Issue: `goimports not installed or available in the PATH`**
- Solution: Run `go install golang.org/x/tools/cmd/goimports@latest` and ensure `$GOPATH/bin` is in your PATH

**Issue: `golangci-lint: command not found`**
- Solution: Run `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest` and ensure `$GOPATH/bin` is in your PATH

**Issue: Pre-commit hooks failing on Windows**
- Solution: Make sure you're using the local hooks configuration (already set up in `.pre-commit-config.yaml`)

## Project Structure

- `cmd/server`: Application entry point.
- `internal/controllers`: Request handlers.
- `internal/routes`: Route definitions.
- `internal/routes/v1`: Version 1 API routes.
- `internal/middleware`: Custom middleware (CORS, logging, etc.).
- `internal/logger`: Centralized logging configuration.
- `docs`: Auto-generated API documentation (Swagger).
