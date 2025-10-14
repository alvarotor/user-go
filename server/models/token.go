package models

import "time"

type Token struct {
	Email               string    `json:"email"`
	Token               string    `json:"token"`
	TokenExpires        time.Time `json:"token_expires"`
	TokenRefresh        string    `json:"refresh_token"`
	TokenRefreshExpires time.Time `json:"refresh_expires"`
}
