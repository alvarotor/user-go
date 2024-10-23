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

type UserResponse struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	Validated     bool   `json:"validated"`
	ProfilePic    string `json:"profilePic"`
	Gender        uint   `json:"gender"`
	Age           uint   `json:"age"`
	CountryOrigin string `json:"countryOrigin"`
}

type UserRequest struct {
	Name          string `json:"name" validate:"required"`
	Validated     bool   `json:"validated"`
	Gender        uint   `json:"gender"`
	Age           uint   `json:"age"`
	CountryOrigin string `json:"countryOrigin"`
}
