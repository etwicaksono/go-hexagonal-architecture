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
	AccessToken      string
	ExpiredAt        time.Time
	RefreshToken     string
	RefreshableUntil time.Time
}

type RefreshTokenRequest struct {
	Token string
}

type AuthToken struct {
	AccessKey string
	ExpiredAt time.Time
	TokenType string
}
