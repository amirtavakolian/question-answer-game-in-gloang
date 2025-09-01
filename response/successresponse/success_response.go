package successresponse

import "QA-Game/response"

type SuccessResponse struct {
	Status  int `json:"status,omitempty"`
	Message string
	Data    interface{} `json:"data"`
}

func NewSuccessResponse() *SuccessResponse {
	return &SuccessResponse{}
}

func (r *SuccessResponse) SetStatus(status int) response.Response {
	r.Status = status
	return r
}

func (r *SuccessResponse) SetMessage(message string) response.Response {
	r.Message = message
	return r
}

func (r *SuccessResponse) SetData(data interface{}) response.Response {
	r.Data = data
	return r
}

func (r *SuccessResponse) GetStatus() int {
	return r.Status
}

func (r *SuccessResponse) Build() response.Response {
	return r
}
