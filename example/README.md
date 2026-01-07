# errx Example (Gin)

This example demonstrates how to use `errx` in a real Gin-based REST API.

---

## Project Structure

```
example/
├── main.go
├── handlers/
│   └── user.go
├── errors/
│   └── user.go
└── go.mod
```

---

## How It Works

1. Errors are defined centrally in `errors/`
2. Business logic returns structured errors
3. Gin middleware converts errors to HTTP responses
4. Client receives stable error codes

---

## Run the Example

```bash
go run .
```

Server will start at:

```
http://localhost:8080
```

---

## API Endpoint

### GET /users/:id

#### Case: User Not Found

```bash
curl http://localhost:8080/users/2
```

**Response:**

```json
{
  "success": false,
  "error": {
    "code": "02-0001",
    "key": "USER_NOT_FOUND",
    "message": "User does not exist"
  }
}
```

#### Case: User Exists

```bash
curl http://localhost:8080/users/1
```

**Response:**

```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "John Cena"
  }
}
```

---

## Key Takeaways

- ✅ Business logic stays clean - only returns domain errors
- ✅ HTTP concerns handled by middleware
- ✅ Consistent error format across all endpoints
- ✅ Stable error codes for client integration
- ✅ Easy to trace and debug issues