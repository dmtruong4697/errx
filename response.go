package errx

type Response struct {
	Success bool       `json:"success"`
	Error   *ErrorBody `json:"error"`
}

type ErrorBody struct {
	Code    string `json:"code"`
	Key     string `json:"key"`
	Message string `json:"message"`
}

func ToResponse(e *Error) Response {
	return Response{
		Success: false,
		Error: &ErrorBody{
			Code:    e.Code,
			Key:     e.Key,
			Message: e.Message,
		},
	}
}
