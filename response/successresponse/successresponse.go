package successresponse

import (
	"encoding/json"
)

type SuccessResponse struct {
	Status  uint
	Message string
	Data    interface{} `json:"data"`
}

func NewSuccessResponse() *SuccessResponse {
	return &SuccessResponse{}
}

func (r *SuccessResponse) SetStatus(status uint) *SuccessResponse {
	r.Status = status
	return r
}

func (r *SuccessResponse) SetMessage(message string) *SuccessResponse {
	r.Message = message
	return r
}

func (r *SuccessResponse) SetData(data interface{}) *SuccessResponse {
	r.Data = data
	return r
}

func (r *SuccessResponse) Buid() string {
	bytes, _ := json.Marshal(r)
	return string(bytes)
}
