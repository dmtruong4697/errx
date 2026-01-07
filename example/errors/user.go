package errors

import "github.com/dmtruong4697/errx"

var UserDomain = errx.NewDomain("USER", "02")

var (
	ErrUserNotFound = UserDomain.New(
		"USER_NOT_FOUND",
		"User does not exist",
		errx.WithHTTPStatus(404),
		errx.WithType(errx.NotFound),
	)

	ErrEmailExists = UserDomain.New(
		"EMAIL_EXISTS",
		"Email already exists",
		errx.WithHTTPStatus(409),
		errx.WithType(errx.Conflict),
	)
)
