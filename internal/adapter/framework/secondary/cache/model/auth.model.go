package model

import "time"

type TokenData struct {
	AccessKey   string    `json:"access_key"`
	ExpiredDate time.Time `json:"expired_date"`
}
