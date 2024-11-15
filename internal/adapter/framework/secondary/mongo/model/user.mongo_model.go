package model

import (
	"github.com/etwicaksono/go-hexagonal-architecture/internal/adapter/core/entity"
	"github.com/guregu/null"
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
	CreatedBy *string            `bson:"created_by"`
	UpdatedAt *time.Time         `bson:"updated_at,omitempty"`
	UpdatedBy *string            `bson:"updated_by,omitempty"`
	DeletedAt *time.Time         `bson:"deleted_at,omitempty"`
	DeletedBy *string            `bson:"deleted_by,omitempty"`
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
		CreatedBy: null.StringFromPtr(u.CreatedBy),
		UpdatedAt: null.TimeFromPtr(u.UpdatedAt),
		UpdatedBy: null.StringFromPtr(u.UpdatedBy),
		DeletedAt: null.TimeFromPtr(u.DeletedAt),
		DeletedBy: null.StringFromPtr(u.DeletedBy),
	}
}

func FromUserEntity(u entity.User) User {
	user := User{
		Email:     u.Email,
		Name:      u.Name,
		Username:  u.Username,
		Password:  u.Password,
		Active:    u.Active,
		CreatedAt: u.CreatedAt,
		CreatedBy: u.CreatedBy.Ptr(),
		UpdatedAt: u.UpdatedAt.Ptr(),
		UpdatedBy: u.UpdatedBy.Ptr(),
		DeletedAt: u.DeletedAt.Ptr(),
		DeletedBy: u.DeletedBy.Ptr(),
	}
	if u.ID != "" {
		user.ID, _ = primitive.ObjectIDFromHex(u.ID)
	}
	return user
}
