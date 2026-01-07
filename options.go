package errx

type Option func(*Error)

func WithHTTPStatus(status int) Option {
	return func(e *Error) {
		e.HTTPStatus = status
	}
}

func WithType(t ErrorType) Option {
	return func(e *Error) {
		e.Type = t
	}
}
