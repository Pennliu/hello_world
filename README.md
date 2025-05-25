[![codecov](https://codecov.io/gh/Pennliu/hello_world/branch/main/graph/badge.svg)](https://codecov.io/gh/Pennliu/hello_world)

# Hello World

A simple Go project demonstrating standard project layout and CI/CD setup.

## Project Structure

```
.
├── cmd/            # Main applications
│   └── hello_world/
├── pkg/            # Library code that's ok to use by external applications
├── internal/       # Private application and library code
├── configs/        # Configuration files
├── scripts/        # Build and deployment scripts
└── test/           # Additional external test applications and test data
```

## Prerequisites

- Go 1.16 or higher

## Getting Started

1. Clone the repository:
```bash
git clone https://github.com/Pennliu/hello_world.git
cd hello_world
```

2. Build the project:
```bash
go build -o hello_world ./cmd/hello_world
```

3. Run the application:
```bash
./hello_world
```

## Development

### Running Tests
```bash
go test ./...
```

### Running Linters
```bash
golangci-lint run
```

## CI/CD

This project uses GitHub Actions for continuous integration and deployment. The workflow is defined in `.github/workflows/ci.yml`.

## License

This project is licensed under the MIT License - see the LICENSE file for details. 