package entity

import "time"

type RegisterRequest struct {
	Email    string
	Name     string
	Username string
	Password string
}

type LoginRequest struct {
	Email    string
	Password string
}

type TokenGenerated struct {
	Token     string
	ExpiredAt time.Time
}
