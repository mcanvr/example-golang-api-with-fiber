# 🚀 Golang Fiber API - DDD Architecture

<div align="center">
<img src="https://golang.org/doc/gopher/doc.png" width="200" height="auto" alt="Go Gopher">
<br>
<strong>Go + Fiber + Domain-Driven Design</strong>
<br>
Sample REST API Project Developed Using Clean Architecture Principles
</div>

## 📋 About the Project

This project is a sample REST API application developed using the Go programming language and Fiber web framework, following Domain-Driven Design (DDD) principles. It can be used as a reference for a modern and scalable API architecture.

## ✨ Key Features

- **📦 Domain-Driven Design**: Clean, layered architecture
- **🔐 JWT Authentication**: Token-based secure API access
- **📚 Swagger Integration**: Complete documentation with OpenAPI
- **🧪 In-Memory Database**: Simple data storage for development
- **⚡ Fiber Web Framework**: High-performance API development
- **🌍 CORS Support**: Cross-Origin Resource Sharing configuration
- **🛡️ Rate Limiting**: Protection against excessive requests
- **⏱️ Request Timeout**: Automatic handling for long-running requests
- **🔍 Validation**: Comprehensive request validation mechanism
- **🌐 Multi-language Support**: Error messages localized in multiple languages

## 🏗️ Project Structure

```
├── cmd
│   └── api                # Application entry points
│       └── main.go
│
├── internal
│   ├── application        # Application services (use cases)
│   │   ├── dto           # Data transfer objects
│   │   └── service       # Application services
│   │
│   ├── domain             # Domain layer
│   │   ├── model         # Domain models (entities, value objects)
│   │   ├── repository    # Repository interfaces
│   │   └── service       # Domain services
│   │
│   ├── infrastructure     # Infrastructure layer
│   │   └── persistence   # Data access implementations
│   │       └── inmemory  # In-memory data storage
│   │
│   ├── interfaces         # Interface layer
│   │   ├── api           # HTTP controllers
│   │   ├── common        # Shared components
│   │   └── middleware    # HTTP middlewares
│   │
│   └── config             # Application configuration
│
├── pkg                    # Shared packages
│   ├── constants          # Constants
│   └── errors             # Custom error types
│
├── docs                   # API documentation
│
└── static                 # Static files (Swagger UI)
```

## 🔑 Core DDD Concepts

The Domain-Driven Design implemented in this project is based on the following principles:

1. **Domain Layer**: The core of business logic. Domain models and services are defined here.
2. **Application Layer**: Orchestrates the domain layer to handle application use cases.
3. **Infrastructure Layer**: Provides implementations for data access and external system interactions.
4. **Interface Layer**: Manages interaction with the outside world (HTTP, CLI, etc.).

## 🛠️ Getting Started

### Requirements

- Go 1.21+
- Git

### Installation

```bash
# Clone the repository
git clone https://github.com/username/example-golang-api-with-fiber.git

# Go to the project directory
cd example-golang-api-with-fiber

# Create necessary folders
make setup

# Install dependencies
go mod download
```

### Environment Variables

The application uses environment variables for configuration. Create a `.env` file in the root directory:

```
SERVER_ADDRESS=:8080
ENVIRONMENT=development
LOG_LEVEL=debug
JWT_SECRET=add_a_strong_secret_key_here
JWT_EXPIRATION_HOURS=24
```

**Note**: A `.env.example` file is provided as a reference.

### Running

```bash
# Run with hot reload (development mode)
make dev

# Run directly
make run

# or
go run cmd/api/main.go
```

The application runs on `http://localhost:8080` by default.

## 📖 API Documentation

The Swagger UI interface can be accessed at:

```
http://localhost:8080/swagger
```

To update the Swagger documentation:

```bash
make swagger
```

## 🔌 API Endpoints

| Method | Endpoint          | Description                        | Auth |
| ------ | ----------------- | ---------------------------------- | ---- |
| POST   | /api/v1/login     | User login and JWT token retrieval | No   |
| GET    | /api/v1/users     | List all users                     | Yes  |
| GET    | /api/v1/users/:id | Get user by ID                     | Yes  |
| POST   | /api/v1/users     | Create new user                    | Yes  |
| PUT    | /api/v1/users/:id | Update user information            | Yes  |
| DELETE | /api/v1/users/:id | Delete user                        | Yes  |

## 🔐 Authentication

A JWT token is required to access protected endpoints. To obtain a token:

```bash
curl -X POST http://localhost:8080/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{"username": "admin", "password": "password"}'
```

To use the token in other requests:

```bash
curl -X GET http://localhost:8080/api/v1/users \
  -H "Authorization: Bearer TOKEN_HERE"
```

## 🔨 Building

```bash
# Build the application
make build

# Clean build outputs and temporary files
make clean
```

## 🛡️ Security Features

### CORS Configuration

The API supports Cross-Origin Resource Sharing (CORS) with configurable settings. By default, requests from local development sources are allowed.

### Rate Limiting

API endpoints are protected against excessive use with rate limiting:

- Default: 50 requests per minute per IP address
- Custom configurations possible through middleware

### Request Timeouts

All requests have a configurable timeout (default: 10 seconds) to prevent hanging connections and resource exhaustion.

## 📝 TODO List

- [x] Add Docker support
- [x] Add CI/CD pipeline
- [x] Add unit tests for domain model
- [ ] Add database integration (PostgreSQL/MySQL)
- [ ] Add more comprehensive tests (integration tests)
- [ ] Implement role-based authorization
- [ ] Performance optimizations
- [ ] Add metrics and monitoring

## 📄 License

MIT
