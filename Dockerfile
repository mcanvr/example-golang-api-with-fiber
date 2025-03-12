# Stage 1: Build the application
FROM golang:1.21-alpine AS builder

# Add git and CA certificates for dependency fetching
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

# Create a non-root user and group
RUN adduser -D -g '' appuser

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first to leverage Docker caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source from the current directory to the working directory inside the container
COPY . .

# Build the application with optimizations
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bin/api cmd/api/main.go

# Stage 2: Create a minimal runtime image
FROM scratch

# Import ca-certificates from builder stage
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# Set timezone
ENV TZ=Europe/Istanbul

# Copy configuration
COPY --from=builder /app/.env.example /.env
COPY --from=builder /app/docs /docs
COPY --from=builder /app/static /static

# Copy the executable from the builder stage
COPY --from=builder /app/bin/api /api

# Use non-root user
USER appuser

# Expose the application port
EXPOSE 8080

# Set environment to production 
ENV ENVIRONMENT=production

# Run the executable
ENTRYPOINT ["/api"] 