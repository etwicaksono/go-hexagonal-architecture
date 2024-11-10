package model

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Email     string             `bson:"email"`
	Name      string             `bson:"name"`
	Username  string             `bson:"username"`
	Password  string             `bson:"password"`
	Active    bool               `bson:"active"`
	CreatedAt time.Time          `bson:"created_at"`
	CreatedBy string             `bson:"created_by"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
	UpdatedBy string             `bson:"updated_by,omitempty"`
	DeletedAt time.Time          `bson:"deleted_at,omitempty"`
	DeletedBy string             `bson:"deleted_by,omitempty"`
}

func (u User) ToEntity() entity.User {
	return entity.User{
		ID:        u.ID.Hex(),
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
	messageItem := User{
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
	if u.ID != "" {
		messageItem.ID, _ = primitive.ObjectIDFromHex(u.ID)
	}
	return messageItem
}
