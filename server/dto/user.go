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
	models.DeviceInfo
}

type ClaimsRefreshResponse struct {
	jwt.RegisteredClaims
	models.DeviceInfo
	CodeRefresh string `json:"codeRefresh"`
}

type UserLogin struct {
	Email string `json:"email" validate:"email,required"`
	models.DeviceInfo
}

type UserUpdate struct {
	Name       string `validate:"required"`
	ProfilePic string
	Bucket     string
}
