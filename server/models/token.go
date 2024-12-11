package models

import "time"

type Token struct {
	Name    string    `json:"name" default:"token"`
	Value   string    `json:"value"`
	Expires time.Time `json:"expires"`
	Email   string    `json:"email"`
}
