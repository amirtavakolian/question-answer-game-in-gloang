package playerparam

type PlayerLoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
