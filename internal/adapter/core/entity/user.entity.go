package entity

import "time"

type User struct {
	ID        string
	Email     string
	Name      string
	Username  string
	Password  string
	Active    bool
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
	DeletedAt time.Time
	DeletedBy string
}

type UserFindFilter struct {
	ID       string
	Email    string
	Name     string
	Username string
	Active   *bool
}

type UserGetFilter struct {
	IDs       []string
	Emails    []string
	Names     []string
	Usernames []string
	Active    *bool
}
