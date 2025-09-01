package richerror

import (
	"QA-Game/response"
	)

type ErrorResponse struct {
	Status  int
	Message string
	Data    interface{}
}

func NewErrorResponse() *ErrorResponse {
	return &ErrorResponse{}
}

func (r *ErrorResponse) SetStatus(status int) response.Response {
	r.Status = status
	return r
}

func (r *ErrorResponse) SetMessage(message string) response.Response {
	r.Message = message
	return r
}

func (r *ErrorResponse) SetData(data any) response.Response {
	r.Data = data
	return r
}

func (r *ErrorResponse) GetStatus() int {
	return r.Status
}

func (r *ErrorResponse) Build() response.Response {

	if r.Data == nil {
		r.Data = ""
	}

	return r
}
