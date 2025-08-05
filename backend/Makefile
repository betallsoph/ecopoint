# ecoPoint Backend Makefile

.PHONY: help proto graphql build run-services docker-build docker-up clean

# Default target
help:
	@echo "Available commands:"
	@echo "  proto       - Generate gRPC code from .proto files"
	@echo "  graphql     - Generate GraphQL resolvers and models"
	@echo "  build       - Build all services"
	@echo "  run-services - Run all services locally"
	@echo "  docker-build - Build Docker images"
	@echo "  docker-up   - Start services with Docker Compose"
	@echo "  clean       - Clean generated files and binaries"

# Install tools
tools:
	go install github.com/99designs/gqlgen@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Generate Protocol Buffer code
proto:
	@echo "Generating gRPC code..."
	mkdir -p proto/generated
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/*.proto

# Generate GraphQL code
graphql:
	@echo "Generating GraphQL code..."
	cd graphql-service && go run github.com/99designs/gqlgen generate

# Build all services
build:
	@echo "Building services..."
	mkdir -p bin
	cd auth-service && go build -o ../bin/auth-service ./main.go
	cd user-service && go build -o ../bin/user-service ./main.go
	cd order-service && go build -o ../bin/order-service ./main.go
	cd gateway && go build -o ../bin/gateway ./main.go
	cd graphql-service && go build -o ../bin/graphql-service ./main.go

# Run services locally
run-services: build
	@echo "Starting MongoDB..."
	docker run -d --name ecopoint-mongo -p 27017:27017 mongo:6.0 || true
	@echo "Starting services..."
	PORT=8081 MONGO_URI=mongodb://localhost:27017/ecopoint JWT_SECRET=test-secret ./bin/auth-service &
	PORT=8082 MONGO_URI=mongodb://localhost:27017/ecopoint JWT_SECRET=test-secret ./bin/user-service &
	PORT=8083 MONGO_URI=mongodb://localhost:27017/ecopoint JWT_SECRET=test-secret ./bin/order-service &
	PORT=4000 MONGO_URI=mongodb://localhost:27017/ecopoint JWT_SECRET=test-secret ./bin/graphql-service &
	PORT=8080 MONGO_URI=mongodb://localhost:27017/ecopoint JWT_SECRET=test-secret ./bin/gateway &
	@echo "Services started. GraphQL Playground: http://localhost:4000"

# Docker commands
docker-build:
	docker-compose build

docker-up:
	docker-compose up -d
	@echo "Services started with Docker. GraphQL Playground: http://localhost:4000"

docker-down:
	docker-compose down

# Clean generated files
clean:
	rm -rf bin/
	rm -rf proto/generated/
	rm -rf graphql-service/generated/
	docker stop ecopoint-mongo 2>/dev/null || true
	docker rm ecopoint-mongo 2>/dev/null || true

# Development helpers
dev-setup: tools proto graphql
	@echo "Development environment setup complete!"

test:
	go test ./...

lint:
	golangci-lint run

# Database helpers
db-seed:
	@echo "Seeding database with sample data..."
	# Add database seeding script here

db-reset:
	docker stop ecopoint-mongo 2>/dev/null || true
	docker rm ecopoint-mongo 2>/dev/null || true
	docker run -d --name ecopoint-mongo -p 27017:27017 mongo:6.0