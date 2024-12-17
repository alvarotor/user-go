package dto

import (
	"github.com/alvarotor/user-go/server/models"
	"github.com/golang-jwt/jwt/v5"
)

type ClaimsResponse struct {
	jwt.RegisteredClaims
	Email      string `json:"email"`
	Admin      bool   `json:"admin"`
	SuperAdmin bool   `json:"superAdmin"`
	models.BaseSecurityLogin
}

type UserLogin struct {
	Email string `json:"email" validate:"email,required"`
	Time  uint   `json:"time" validate:"required,gt=0,lt=366"`
	models.BaseSecurityLogin
}

type UserUpdate struct {
	Name       string `validate:"required"`
	ProfilePic string
	Bucket     string
}
