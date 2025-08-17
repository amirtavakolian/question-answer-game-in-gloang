package successresponse

import "fmt"

type SuccessResponse struct {
	status  uint
	message string
	data    any
}

func NewSuccessResponse() *SuccessResponse {
	return &SuccessResponse{}
}

func (r *SuccessResponse) SetStatus(status uint) *SuccessResponse {
	r.status = status
	return r
}

func (r *SuccessResponse) SetMessage(message string) *SuccessResponse {
	r.message = message
	return r
}

func (r *SuccessResponse) SetData(data any) *SuccessResponse {
	r.data = data
	return r
}

func (r *SuccessResponse) Buid() string {
	return fmt.Sprintf(`{
		"status": "%d",
		"message": "%s",
		"data": "%v"
	}`, r.status, r.message, r.data)
}
