package models

import "time"

type Token struct {
	Name           string    `json:"name" default:"token"`
	Token          string    `json:"value"`
	ExpiresRefresh time.Time `json:"expires_refresh"`
	Email          string    `json:"email"`
	RefreshToken   string    `json:"refresh_token"`
}
