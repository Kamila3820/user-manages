
# 🏢 User Management
A modern manage user backend service built in Go with high-performance APIs using **gRPC**, **Echo**, **gRPC-Server**, and **MongoDB** to manage user data. This project follows a clean 3-layer architecture: **Controller → Service → Repository** for better modularity and maintainability. The project also features secure authentication, and credential management.

---

### 📂 Project Structure (Simplified)

```

.

├── *config*/         # Configuration files (e.g., MongoDB setup)
├── *env*/            # .env and environment management
├── *modules*/        # Modules for business logic (user module, etc.)
├── *pkg*/            # Shared utilities and core packages
├── *server*/         # HTTP server and route registration
├── *test*/           # Test files and sample data
├── *go.mod / sum*    # Go modules
└── *README.md*       # Project documentation
````

---

## 🚀 Features

### ✅ Core Functionality

- **User Account Management** (Registration, Login, Logout)
- **User Data Management** (Find User By Id, List Users, Update User Name/Email, Delete User)
- **Secure Authentication**
  - Access & Refresh Token & API Key (JWT)
  - Token revocation and expiry handling
- **Dockerized MongoDB support** 

### 🔁 gRPC + REST

- **gRPC** for fast, type-safe RPCs
- **gRPC-Server** provides RESTful HTTP APIs
- **Echo** as a lightweight, simple to use with minimal boilerplate

### 🔐 Security

- Secure password hashing 
- Access token stored in headers; refresh token managed securely

#### ✅ Token Types
- Access Token: Short-lived token for authenticating user requests
- Refresh Token: Longer-lived token used to obtain new access tokens
- API Key: Long-lived token for internal service communication

#### 🛠 Token Generation
Tokens are generated using the **SignToken()** method from the **AuthFactory** interface:

``` bash
AccessToken(cfg *config.Config, claims *jwtauth.Claims) string
RefreshToken(cfg *config.Config, claims *jwtauth.Claims) string
```
Internally uses:
``` bash
jwtauth.NewAccessToken(secret, expiry, claims)
jwtauth.NewRefreshToken(secret, expiry, claims)
```

#### 🔄 Token Refresh Logic

1. Client sends refresh token

2. FindOneUserProfileToRefresh() verifies the token via gRPC

3. New access/refresh tokens generated

4. Old credentials updated using UpdateOneUserCredential()

#### 🗃 Example of the Token Storage
``` bash
{
  "_id": ObjectId,
  "user_id": "string",
  "access_token": "string",
  "refresh_token": "string",
  "created_at": "timestamp",
  "updated_at": "timestamp"
}
```

#### 🔑 JWT API Key Setup
The server sets an API key for internal authorization in gRPC calls:
``` bash
func SetApiKeyInContext(pctx *context.Context) {
	*pctx = metadata.NewOutgoingContext(*pctx, metadata.Pairs("auth", apiKeyInstant))
}
```

### 🛡️ Middleware & Server Lifecycle
- **Request Timeout** (Limits the maximum duration for each request to prevent hanging requests)
- **CORS** (Allows requests from any origin and supports common HTTP methods)
- **Body Limit** (estricts the size of incoming request bodies to 10 megabytes.)
- **Logger Middleware**  (Logs each HTTP request for monitoring and debugging.)


## 🧠 Getting Started

### 📦 Packages
```bash
go get github.com/labstack/echo/v4
go get github.com/labstack/echo/v4/middleware
go get github.com/go-playground/validator/v10
go get github.com/joho/godotenv
go get go.mongodb.org/mongo-driver/mongo
go get github.com/golang-jwt/jwt/v5
go get github.com/stretchr/testify
go get github.com/IBM/sarama
```

### 📃 Start App in Terminal
```bash
go run main.go ./env/dev/.env.auth
go run main.go ./env/dev/.env.user
```

### 🍃 MongoDb
Start Docker Compose

```bash
docker compose -f docker-compose.db.yml up -d
```

Enter into a db container

```bash
docker exec -it <db_name> bash
```

### Migration

dev

```bash
go run ./pkg/database/script/migration.go ./env/dev/.env.user && \
go run ./pkg/database/script/migration.go ./env/dev/.env.auth && \
```

### Generate a Proto File Command

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



### 🧪 Testing
Written using Go’s testing package

MongoDB interactions mocked via interfaces

### 📦 Sample of API requests/responses
- Endpoint: ``` POST /auth_v1/auth/login ```

**StatusOK Response (200)**:
``` bash
{
  "_id": "user:6862c93ba8657cc777fca79b",
  "email": "fang@example.com",
  "name": "Fang Name",
  "created_at": "2025-07-01T00:28:27.027+07:00",
  "credential": {
    "_id": "6862c9a42a5301fb1bc008c3",
    "user_id": "user:6862c93ba8657cc777fca79b",
    "access_token": "<JWT_ACCESS_TOKEN>",
    "refresh_token": "<JWT_REFRESH_TOKEN>",
    "created_at": "2025-07-01T00:30:12.179+07:00",
    "updated_at": "2025-07-01T00:30:12.179+07:00"
  }
}
```

**Unauthorized Response (401)**:
``` bash
{
    "message": "error: email or password is incorrect"
}
```

- Endpoint: ``` GET /user_v1/user/user:xxxxxxxxxx ```
**StatusOK Response (200)**:
``` bash
{
    "_id": "6862a742e166ea321970323f",
    "email": "user002@sekai.com",
    "name": "User002",
    "created_at": "2025-06-30T22:03:30.066+07:00"
}
```

**Bad Request Response (400)**:
``` bash
{
    "message": "error: user profile not found"
}
```