package entity

import "time"

type UserAccess struct {
	Id        int64
	UserId    int64
	Key       string
	Platform  string
	UserAgent string
	ExpiredAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
