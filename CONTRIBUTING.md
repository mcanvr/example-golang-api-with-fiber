# Contributing Guide

Thank you for your interest in contributing to this project! This document provides guidelines for those who want to contribute to the project.

## Development Environment Setup

1. Fork the repository and clone it to your local machine
2. Install the necessary dependencies: `go mod download`
3. Start the development server: `make dev`

## Code Standards

- Follow the official Go code format and standards
- Run `go fmt` and `go vet` before writing code
- Add documentation in godoc format for all public functions and types
- Ensure your code has high test coverage

## Pull Request Process

1. Create a new branch for your changes (`git checkout -b feature/amazing-feature`)
2. Commit your changes (`git commit -m 'feat: added amazing feature'`)
3. Push your branch (`git push origin feature/amazing-feature`)
4. Open a Pull Request on GitHub

## Commit Messages

Use the [Conventional Commits](https://www.conventionalcommits.org/) format for your commit messages:

- `feat`: A new feature
- `fix`: A bug fix
- `docs`: Documentation only changes
- `style`: Changes that do not affect the meaning of the code (formatting, etc.)
- `refactor`: A code change that neither fixes a bug nor adds a feature
- `perf`: A code change that improves performance
- `test`: Adding missing tests or correcting existing tests
- `chore`: Changes to the build process or auxiliary tools

For example:

```
feat: added user profile interface
fix: fixed authentication token duration
docs: updated README file
```

## Testing

- Submit new features and bug fixes with appropriate unit tests
- Ensure all tests pass: `go test ./...`

## Bug Reports

When you find a bug, please create a report in the GitHub Issues section with the following information:

- A brief summary of the bug and the expected behavior
- Steps to reproduce the bug
- Your working environment (operating system, Go version, etc.)
- Screenshots or error logs if available

## Feature Requests

To propose a new feature, create a request in the GitHub Issues section and make sure it includes:

- A clear description of the feature
- An explanation of why this feature is needed
- Possible implementation methods or sample code (if available)

## License

By contributing to this project, you agree that your contributions will be licensed under the project's [MIT License](LICENSE).
