package richerror

import "fmt"

type ErrorResponse struct {
	status  uint
	message string
	data    any
}

func NewErrorResponse() *ErrorResponse {
	return &ErrorResponse{}
}

func (r *ErrorResponse) SetStatus(status uint) *ErrorResponse {
	r.status = status
	return r
}

func (r *ErrorResponse) SetMessage(message string) *ErrorResponse {
	r.message = message
	return r
}

func (r *ErrorResponse) SetData(data any) *ErrorResponse {
	r.data = data
	return r
}

func (r *ErrorResponse) Buid() string {

	if r.data == nil {
		r.data = ""
	}

	return fmt.Sprintf(`{
		"status": "%d",
		"message": "%s",
		"data": "%v"
	}`, r.status, r.message, r.data)
}
