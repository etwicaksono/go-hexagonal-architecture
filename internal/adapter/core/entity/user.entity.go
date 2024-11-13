package entity

import (
	"github.com/guregu/null"
	"time"
)

type User struct {
	ID        string
	Email     string
	Name      string
	Username  string
	Password  string
	Active    bool
	CreatedAt time.Time
	CreatedBy null.String
	UpdatedAt null.Time
	UpdatedBy null.String
	DeletedAt null.Time
	DeletedBy null.String
}

type UserFindFilter struct {
	ID       null.String
	Email    null.String
	Name     null.String
	Username null.String
	Active   null.Bool
}

type UserGetFilter struct {
	IDs       []string
	Emails    []string
	Names     []string
	Usernames []string
	Active    null.Bool
}
