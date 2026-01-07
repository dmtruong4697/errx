package errx

func Wrap(e *Error) *Error {
	return &Error{
		Code:       e.Code,
		Key:        e.Key,
		Message:    e.Message,
		HTTPStatus: e.HTTPStatus,
		Type:       e.Type,
		Meta:       map[string]any{},
	}
}

func (e *Error) WithMeta(key string, val any) *Error {
	if e.Meta == nil {
		e.Meta = map[string]any{}
	}
	e.Meta[key] = val
	return e
}

func (e *Error) WithCause(err error) *Error {
	e.Cause = err
	return e
}
