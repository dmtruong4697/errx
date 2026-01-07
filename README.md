# errx

`errx` is a lightweight Go library for **standardizing API errors** with:
- Structured error codes
- HTTP status mapping
- Centralized error definitions
- Clean integration with REST frameworks (Gin first-class)

---

## Why errx?

Common problems in REST APIs:
- Error codes are duplicated or inconsistent
- Error messages are hardcoded everywhere
- Client cannot distinguish error cases reliably
- Support team cannot trace issues efficiently

`errx` solves this by enforcing:
- Centralized error definitions
- Unique, traceable error codes
- Clear separation between business logic and HTTP response

---

## Features

- ✅ Unique error codes per domain
- ✅ Auto-increment error code generation
- ✅ HTTP status mapping
- ✅ Error wrapping with metadata
- ✅ Gin middleware support
- ✅ Framework-agnostic core

---

## Installation

```bash
go get github.com/dmtruong4697/errx
```

---

## Basic Concepts

### Error Structure

```json
{
  "code": "02-0001",
  "key": "USER_NOT_FOUND",
  "message": "User does not exist"
}
```

Each error contains:
- **`code`**: stable error code for client & support
- **`key`**: logical identifier
- **`message`**: default human-readable message

### Define Errors (Centralized)

```go
package errors

import "github.com/dmtruong4697/errx"

var UserDomain = errx.NewDomain("USER", "02")

var ErrUserNotFound = UserDomain.New(
	"USER_NOT_FOUND",
	"User does not exist",
	errx.WithHTTPStatus(404),
	errx.WithType(errx.NotFound),
)
```

### Use in Business Logic

```go
if user == nil {
	return errx.Wrap(ErrUserNotFound).
		WithMeta("user_id", id)
}
```

Business logic:
- Does NOT return HTTP response
- Only returns structured errors

### Gin Middleware

```go
r.Use(middleware.ErrorHandler())
```

The middleware:
- Catches all errors
- Converts them to standardized JSON responses
- Preserves HTTP status codes

### Example Response

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

---

## Error Wrapping

```go
errx.Wrap(ErrUserNotFound).
	WithMeta("user_id", 123).
	WithCause(dbErr)
```

Metadata is useful for logging and tracing.

---

## Roadmap

- [ ] i18n support
- [ ] Error documentation generator
- [ ] `errors.Is` / `errors.As` compatibility
- [ ] OpenTelemetry integration

---

## Philosophy

- Errors are API contracts
- Error codes should be stable
- Error handling should be boring and predictable

---

## License

MIT