dev:
	air 

run:
	go run cmd/api/main.go

swagger:
	swag init -g cmd/api/main.go -pd -output docs
	
build:
	go build -o ./bin/api ./cmd/api/main.go

clean:
	rm -rf ./bin
	rm -rf ./tmp

setup:
	mkdir -p bin
	mkdir -p tmp

docker-build:
	docker build -t golang-fiber-api:latest .

docker-run:
	docker run -p 8080:8080 --name golang-fiber-api -d golang-fiber-api:latest

docker-stop:
	docker stop golang-fiber-api
	docker rm golang-fiber-api

docker-compose-up:
	docker-compose up -d

docker-compose-down:
	docker-compose down

lint:
	go fmt ./...
	go vet ./...
	staticcheck ./...

test:
	go test -v ./...

.PHONY: dev swagger run build clean setup docker-build docker-run docker-stop docker-compose-up docker-compose-down lint test 