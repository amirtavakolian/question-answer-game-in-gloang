package response

type Response interface {
	SetStatus(status int) Response
	SetMessage(message string) Response
	SetData(data any) Response
	Build() Response
	GetStatus() int
}
