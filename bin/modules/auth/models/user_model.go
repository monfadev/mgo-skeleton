package models

import "time"

type UserModel struct {
	ID        int
	Name      string
	Email     string
	Role      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
