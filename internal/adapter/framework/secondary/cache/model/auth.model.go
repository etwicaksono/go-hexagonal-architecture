package model

import "time"

type TokenData struct {
	Token        string    `json:"token_code"`
	RefreshToken string    `json:"refresh_token"`
	ExpiredDate  time.Time `json:"expired_date"`
}
