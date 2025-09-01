package entity

import "time"

type Permission struct {
	Id          uint
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
