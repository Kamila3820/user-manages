# 🏢 User Management
A modern manage user backend service built in Go with high-performance APIs using **gRPC**, **Echo**, **gRPC-Server**, and **MongoDB** to manage user data. This project follows a clean 3-layer architecture: **Controller → Service → Repository** for better modularity and maintainability. The project also features secure authentication, and credential management.

## 📁 Project Structure

├── config/         # Configuration files (e.g., MongoDB setup)
├── env/            # .env and environment management
├── modules/        # Modules for business logic (user module, etc.)
├── pkg/            # Shared utilities and core packages
├── server/         # HTTP server and route registration
├── test/           # Test files and sample data
├── go.mod / sum    # Go modules
└── README.md       # Project documentation

## 🚀 Features
- List all users

- Register new users

- Update a user’s name/email

- Delete a user

- JWT auth ready (access, refresh, and API tokens)

- gRPC-ready architecture (optional)

- Well-separated request/response models

- ockerized MongoDB support

## 📦 Packages
go get github.com/labstack/echo/v4
go get github.com/labstack/echo/v4/middleware
go get github.com/go-playground/validator/v10
go get github.com/joho/godotenv
go get go.mongodb.org/mongo-driver/mongo
go get github.com/golang-jwt/jwt/v5
go get github.com/stretchr/testify
go get github.com/IBM/sarama

## 📃 Start App in Terminal
```bash
go run main.go ./env/dev/.env.auth
go run main.go ./env/dev/.env.user
```

## 🍃 MongoDb
Start Docker Compose

```bash
docker compose -f docker-compose.db.yml up -d
```

Enter into a db container

```bash
docker exec -it <db_name> bash
```

## Migration

dev

```bash
go run ./pkg/database/script/migration.go ./env/dev/.env.user && \
go run ./pkg/database/script/migration.go ./env/dev/.env.auth && \
```

## Generate a Proto File Command

player

```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    modules/user/userPb/userPb.proto
```

auth

```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    modules/auth/authPb/authPb.proto
```

## 🧪 Testing
Written using Go’s testing package

MongoDB interactions mocked via interfaces
