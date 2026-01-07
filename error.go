package errx

type ErrorType string

const (
	Validation ErrorType = "VALIDATION"
	NotFound   ErrorType = "NOT_FOUND"
	Conflict   ErrorType = "CONFLICT"
	Internal   ErrorType = "INTERNAL"
)

type Error struct {
	Code       string
	Key        string
	Message    string
	HTTPStatus int
	Type       ErrorType
	Meta       map[string]any
	Cause      error
}

func (e *Error) Error() string {
	return e.Message
}
