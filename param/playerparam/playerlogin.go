package playerparam

type PlayerLoginRequest struct {
	PlayerId      int
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
