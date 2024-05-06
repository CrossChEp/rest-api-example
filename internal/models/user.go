package models

import "time"

type User struct {
	ID        UserID
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	DeletedAt time.Time
}

type UserFilter struct {
	IDs       []UserID
	Names     []string
	Emails    []string
	Passwords []string
}
