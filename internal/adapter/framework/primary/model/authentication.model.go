package model

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
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
