package playerdto

import "time"

type PlayerRegister struct {
	Name        string    `json:"first_name"`
	PhoneNumber string    `json:"phone_number"`
	Password    string    `json:"password"`
	Avatar      string    `json:"avatar,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
