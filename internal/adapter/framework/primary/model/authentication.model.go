package model

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
)

type RegisterRequest struct {
	Email    string `json:"email" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (r RegisterRequest) ToEntity() entity.RegisterRequest {
	return entity.RegisterRequest(r)
}

func FromRegisterRequestEntity(r entity.RegisterRequest) RegisterRequest {
	return RegisterRequest(r)
}
