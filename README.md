# Sens.ai Backend

Go REST API project using Fiber.

## Prerequisites

- Go 1.20+
- [golangci-lint](https://golangci-lint.run/usage/install/)
- [pre-commit](https://pre-commit.com/#install)

## Setup

1. Clone the repository.
2. Install dependencies:
   ```bash
   make deps
   ```
3. Install pre-commit hooks:
   ```bash
   pre-commit install
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

## Project Structure

- `cmd/server`: Application entry point.
- `internal/controllers`: Request handlers.
- `internal/routes`: Route definitions.
- `internal/routes/v1`: Version 1 API routes.
