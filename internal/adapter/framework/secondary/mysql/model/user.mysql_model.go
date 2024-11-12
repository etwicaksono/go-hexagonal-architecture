package model

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        string    `gorm:"column:id;type:char(36);primaryKey"`
	Email     string    `gorm:"column:email;type:varchar(64)"`
	Name      string    `gorm:"column:name;type:varchar(64)"`
	Username  string    `gorm:"column:username;type:varchar(64)"`
	Password  string    `gorm:"column:password;type:varchar(255)"`
	Active    bool      `gorm:"column:active;type:tinyint(1);default:1"`
	CreatedAt time.Time `gorm:"column:created_at;"`
	CreatedBy string    `gorm:"column:created_by;type:varchar(255)"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:NULL"`
	UpdatedBy string    `gorm:"column:updated_by;type:varchar(255)"`
	DeletedAt time.Time `gorm:"column:deleted_at;default:NULL"`
	DeletedBy string    `gorm:"column:deleted_by;type:varchar(255)"`
}

func (u User) ToEntity() entity.User {
	userId := u.ID
	if u.ID == "" {
		userId = uuid.New().String()
	}
	return entity.User{
		ID:        userId,
		Email:     u.Email,
		Name:      u.Name,
		Username:  u.Username,
		Password:  u.Password,
		Active:    u.Active,
		CreatedAt: u.CreatedAt,
		CreatedBy: u.CreatedBy,
		UpdatedAt: u.UpdatedAt,
		UpdatedBy: u.UpdatedBy,
		DeletedAt: u.DeletedAt,
		DeletedBy: u.DeletedBy,
	}
}

func FromUserEntity(u entity.User) User {
	userId := u.ID
	if u.ID == "" {
		userId = uuid.New().String()
	}
	user := User{
		ID:        userId,
		Email:     u.Email,
		Name:      u.Name,
		Username:  u.Username,
		Password:  u.Password,
		Active:    u.Active,
		CreatedAt: u.CreatedAt,
		CreatedBy: u.CreatedBy,
		UpdatedAt: u.UpdatedAt,
		UpdatedBy: u.UpdatedBy,
		DeletedAt: u.DeletedAt,
		DeletedBy: u.DeletedBy,
	}
	return user
}
