package entity

import "time"

type Player struct {
	Id          uint
	Name        string
	PhoneNumber string
	Password    string `json:"-"`
	Avatar      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
