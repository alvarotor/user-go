package controllers

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/alvarotor/user-go/server/dto"
	"github.com/alvarotor/user-go/server/models"
	"github.com/golang-jwt/jwt/v5"
)

func (u *controllerUser) Validate(c context.Context, code string) (int, models.Token, error) {
	user, err := u.GetByCode(c, code)
	if err != nil {
		u.log.Error(err.Error())
		return http.StatusNotFound, models.Token{}, err
	}

	if user == nil {
		u.log.Error(models.ErrInvalidCode.Error())
		return http.StatusBadRequest, models.Token{}, models.ErrInvalidCode
	}

	if user.Code == "OUT" || strings.TrimSpace(user.Code) == "" {
		u.log.Error(models.ErrInvalidCode.Error())
		return http.StatusBadRequest, models.Token{}, models.ErrInvalidCode
	}

	if user.CodeExpire.Before(time.Now().UTC()) {
		u.log.Error(models.ErrExpiredCode.Error())
		return http.StatusBadRequest, models.Token{}, models.ErrExpiredCode
	}

	if !user.Validated {
		err = u.ValidateSvc(c, user.Email)
		if err != nil {
			u.log.Error(err.Error())
			return http.StatusInternalServerError, models.Token{}, err
		}
	}

	expirationTime := getExpirationTime(uint(u.conf.TokenExpirationTime))
	u.log.Info("Generating access token", "email", user.Email, "token_expiration_config", u.conf.TokenExpirationTime, "expires_at", expirationTime.Unix(), "now", time.Now().UTC().Unix())

	claims := &dto.ClaimsResponse{
		Email:      user.Email,
		Admin:      user.Admin,
		SuperAdmin: user.SuperAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    u.conf.Issuer,
		},
		DeviceInfo: createDeviceInfo(user),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(u.conf.JWTKey)
	if err != nil {
		u.log.Error(err.Error())
		return http.StatusInternalServerError, models.Token{}, err
	}

	expirationTimeRefresh := getExpirationTime(uint(u.conf.TokenExpirationTimeRefresh))
	codeRefreshPreview := user.CodeRefresh
	if len(codeRefreshPreview) > 8 {
		codeRefreshPreview = codeRefreshPreview[:8] + "..."
	}
	u.log.Info("Generating refresh token", "email", user.Email, "refresh_expiration_config", u.conf.TokenExpirationTimeRefresh, "code_refresh", codeRefreshPreview, "refresh_expires_at", expirationTimeRefresh.Unix(), "now", time.Now().UTC().Unix())

	claimsRefresh := &dto.ClaimsRefreshResponse{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTimeRefresh),
			Issuer:    u.conf.Issuer,
		},
		CodeRefresh: user.CodeRefresh,
		DeviceInfo:  createDeviceInfo(user),
	}
	tokenRefresh := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	tokenRefreshString, err := tokenRefresh.SignedString(u.conf.JWTKey)
	if err != nil {
		u.log.Error(err.Error())
		return http.StatusInternalServerError, models.Token{}, err
	}

	model := models.Token{
		Email:               user.Email,
		Token:               tokenString,
		TokenExpires:        expirationTime,
		TokenRefresh:        tokenRefreshString,
		TokenRefreshExpires: expirationTimeRefresh,
	}

	u.log.Info("Validation completed successfully",
		"email", user.Email,
		"token_length", len(tokenString),
		"refresh_token_length", len(tokenRefreshString),
		"expires_in_seconds", int(expirationTime.Sub(time.Now().UTC()).Seconds()),
		"refresh_expires_in_seconds", int(expirationTimeRefresh.Sub(time.Now().UTC()).Seconds()))

	return http.StatusOK, model, nil
}

func getExpirationTime(seconds uint) time.Time {
	var expirationTime time.Time
	now := time.Now().UTC()
	expirationTime = now.Add(time.Duration(seconds) * time.Second)
	return expirationTime
}

func createDeviceInfo(user *models.User) models.DeviceInfo {
	return models.DeviceInfo{
		Browser:                user.Browser,
		BrowserVersion:         user.BrowserVersion,
		OperatingSystem:        user.OperatingSystem,
		OperatingSystemVersion: user.OperatingSystemVersion,
		Cpu:                    user.Cpu,
		Language:               user.Language,
		Timezone:               user.Timezone,
		CookiesEnabled:         user.CookiesEnabled,
	}
}
