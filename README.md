![Build Status](https://github.com/Pennliu/hello_world/actions/workflows/ci.yml/badge.svg)
[![codecov](https://codecov.io/gh/Pennliu/hello_world/branch/main/graph/badge.svg)](https://codecov.io/gh/Pennliu/hello_world)
[![Go Report Card](https://goreportcard.com/badge/github.com/Pennliu/hello_world)](https://goreportcard.com/report/github.com/Pennliu/hello_world)
![Go Version](https://img.shields.io/badge/go-1.24+-blue)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![GitHub release](https://img.shields.io/github/v/release/Pennliu/hello_world.svg)](https://github.com/Pennliu/hello_world/releases)
[![Issues](https://img.shields.io/github/issues/Pennliu/hello_world.svg)](https://github.com/Pennliu/hello_world/issues)
[![PRs](https://img.shields.io/github/issues-pr/Pennliu/hello_world.svg)](https://github.com/Pennliu/hello_world/pulls)
[![GitHub stars](https://img.shields.io/github/stars/Pennliu/hello_world.svg)](https://github.com/Pennliu/hello_world/stargazers)

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

- Go 1.24 or higher ([Download Go](https://go.dev/dl/))

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