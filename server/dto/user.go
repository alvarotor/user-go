package dto

import (
	"github.com/golang-jwt/jwt/v5"
)

type ClaimsResponse struct {
	jwt.RegisteredClaims
	ID         uint   `json:"id"`
	Email      string `json:"email"`
	Admin      bool   `json:"admin"`
	SuperAdmin bool   `json:"superAdmin"`
}

type UserLogin struct {
	Email string `json:"email" validate:"email,required"`
	Time  uint   `json:"time" validate:"required,gt=0 lt=60"`
}
