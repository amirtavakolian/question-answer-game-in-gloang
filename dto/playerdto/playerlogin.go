package playerdto

type PlayerLogin struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
