package errx

import "fmt"

type Domain struct {
	Name     string
	Prefix   string
	sequence int
}

func NewDomain(name, prefix string) *Domain {
	return &Domain{
		Name:   name,
		Prefix: prefix,
	}
}

func (d *Domain) New(
	key string,
	message string,
	opts ...Option,
) *Error {

	d.sequence++
	code := fmt.Sprintf("%s-%04d", d.Prefix, d.sequence)

	register(code)

	e := &Error{
		Code:       code,
		Key:        key,
		Message:    message,
		HTTPStatus: 500,
		Type:       Internal,
	}

	for _, opt := range opts {
		opt(e)
	}

	return e
}
