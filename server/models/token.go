package models

import "time"

type Token struct {
	Name           string    `json:"name" default:"token"`
	Token          string    `json:"value"`
	Expires        time.Time `json:"expires"`
	Email          string    `json:"email"`
	RefreshToken   string    `json:"refresh_token"`
}
