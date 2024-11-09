package model

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"time"
)

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required,is-username"`
	Password string `json:"password" validate:"required,max=72"`
}

func (r RegisterRequest) ToEntity() entity.RegisterRequest {
	return entity.RegisterRequest(r)
}

func FromRegisterRequestEntity(r entity.RegisterRequest) RegisterRequest {
	return RegisterRequest(r)
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,max=72"`
}

func (r LoginRequest) ToEntity() entity.LoginRequest {
	return entity.LoginRequest(r)
}

func FromLoginRequestEntity(r entity.LoginRequest) LoginRequest {
	return LoginRequest(r)
}

type TokenGenerated struct {
	AccessToken      string    `json:"access_token"`
	ExpiredAt        time.Time `json:"expired_at"`
	RefreshToken     string    `json:"refresh_token"`
	RefreshableUntil time.Time `json:"refreshable_until"`
}

func (t TokenGenerated) ToEntity() entity.TokenGenerated {
	return entity.TokenGenerated(t)
}

func FromTokenGeneratedEntity(t entity.TokenGenerated) TokenGenerated {
	return TokenGenerated(t)
}

type TokenReversed struct {
	AccessKey string
	ExpiredAt time.Time
	TokenType string
}

type RefreshTokenRequest struct {
	Token string `json:"refresh_token" validate:"required"`
}
