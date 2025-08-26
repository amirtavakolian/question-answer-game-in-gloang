package dbresponses

type Player struct {
	UserId      int
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
