package entity

import "time"

type User struct {
	ID        string
	Email     string
	Name      string
	Username  string
	Password  string
	Active    bool
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
	DeletedAt *time.Time
	DeletedBy string
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
