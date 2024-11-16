package model

import "time"

type AuthCachedData struct {
	UserId    string    `json:"user_id"`
	AccessKey string    `json:"access_key"`
	ExpiredAt time.Time `json:"expired_at"`
}
