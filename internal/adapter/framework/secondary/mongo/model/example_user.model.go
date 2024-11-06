package model

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Email     string             `bson:"Email"`
	Name      string             `bson:"Name"`
	Username  string             `bson:"Username"`
	Password  string             `bson:"Password"`
	Active    bool               `bson:"Active"`
	CreatedAt time.Time          `bson:"CreatedAt"`
	CreatedBy string             `bson:"CreatedBy"`
	UpdatedAt time.Time          `bson:"UpdatedAt"`
	UpdatedBy string             `bson:"UpdatedBy"`
	DeletedAt *time.Time         `bson:"DeletedAt"`
	DeletedBy string             `bson:"DeletedBy"`
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
