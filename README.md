# User and Product API

Simple REST API built with [Gin](https://github.com/gin-gonic/gin) providing JWT authentication and CRUD operations for users and products.

## Prerequisites

- Go 1.21 or later

## Getting Started

1. Install dependencies:

   ```bash
   go mod download
   ```

2. Run the server:

   ```bash
   go run main.go
   ```

3. Open your browser at `http://localhost:8080/docs/index.html` for the Swagger UI.

## API Endpoints

| Method | Path | Description | Auth |
| ------ | ---- | ----------- | ---- |
| POST   | `/register` | Register a new user | None |
| POST   | `/login` | Obtain JWT token | None |
| GET    | `/users` | List users | Bearer |
| GET    | `/users/{id}` | Get user by ID | Bearer |
| POST   | `/users` | Create user | Bearer |
| PUT    | `/users/{id}` | Update user | Bearer |
| DELETE | `/users/{id}` | Delete user | Bearer |
| GET    | `/products` | List products | Bearer |
| GET    | `/products/{id}` | Get product by ID | Bearer |
| POST   | `/products` | Create product | Bearer |
| PUT    | `/products/{id}` | Update product | Bearer |
| DELETE | `/products/{id}` | Delete product | Bearer |

## Development

Swagger documentation is generated using [swag](https://github.com/swaggo/swag).
To regenerate the docs after making changes:

```bash
swag init -g main.go
```

