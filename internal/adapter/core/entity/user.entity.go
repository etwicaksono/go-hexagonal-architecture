package entity

import "time"

type User struct {
	Id        string
	Username  string
	Phone     string
	Email     string
	FullName  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type RegisterRequest struct {
	Username string
	Phone    string
	Email    string
	FullName string
	Password string
}

type LoginRequest struct {
	Username string
	Password string
}
