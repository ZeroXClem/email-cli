# Contributing Guidelines

## Getting Started

1. Fork the repository
2. Clone your fork: `git clone https://github.com/yourusername/email-cli`
3. Create a new branch: `git checkout -b feature/your-feature-name`

## Development Setup

1. Install Go (1.16 or later)
2. Install dependencies: `go mod tidy`
3. Run tests: `go test ./...`

## Code Style

- Follow standard Go formatting: `go fmt ./...`
- Use `golint` and `go vet` for code quality
- Write descriptive commit messages
- Add tests for new features

## Pull Request Process

1. Update documentation for new features
2. Add or update tests as needed
3. Ensure all tests pass
4. Update the README.md if needed
5. Create a pull request with a clear description

## Testing

- Write unit tests for new packages
- Add integration tests for new features
- Test edge cases and error conditions

## Documentation

- Update usage.md for new features
- Add command-line flag documentation
- Include example configurations
- Document any breaking changes

## Code Review

- All submissions require review
- Address review comments promptly
- Keep pull requests focused and atomic

## License

By contributing, you agree that your contributions will be licensed under the project's license.