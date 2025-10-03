package controllers

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/alvarotor/user-go/server/dto"
	"github.com/alvarotor/user-go/server/models"
	"github.com/golang-jwt/jwt/v5"
)

func (u *controllerUser) TokenToUser(
	c context.Context,
	token string,
	browser string,
	browserVersion string,
	operatingSystem string,
	operatingSystemVersion string,
	cpu string,
	language string,
	timezone string,
	cookiesEnabled bool,
) (*models.User, error) {
	u.log.Info("TokenToUser validation started", "token_length", len(token))

	claims := &dto.ClaimsResponse{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (any, error) {
		return u.conf.JWTKey, nil
	})

	u.log.Info("JWT parsing completed", "email", claims.Email, "exp", claims.ExpiresAt, "now", time.Now().UTC())

	if err := u.validateToken(tkn, err); err != nil {
		u.log.Error("Token validation failed", "error", err.Error())
		return &models.User{}, err
	}

	u.log.Info("Token validation successful", "email", claims.Email)

	user, err := u.GetByEmail(c, claims.Email)
	if err != nil {
		return &models.User{}, err
	}
	if user.Code == "OUT" {
		u.log.Error(models.ErrUserNotLogged.Error())
		return &models.User{}, models.ErrUserNotLogged
	}
	if len(user.Code) != u.conf.SizeRandomStringValidation {
		u.log.Error(models.ErrInvalidUser.Error())
		return &models.User{}, models.ErrInvalidUser
	}

	secs := models.DeviceInfo{
		Browser:                browser,
		BrowserVersion:         browserVersion,
		OperatingSystem:        operatingSystem,
		OperatingSystemVersion: operatingSystemVersion,
		Cpu:                    cpu,
		Language:               language,
		Timezone:               timezone,
		CookiesEnabled:         cookiesEnabled,
	}

	// More robust device validation - allow for minor differences
	if !u.validateDeviceInfo(claims.DeviceInfo, secs) {
		u.log.Error("Device validation failed",
			"email", claims.Email,
			"jwt_browser", claims.DeviceInfo.Browser,
			"jwt_os", claims.DeviceInfo.OperatingSystem,
			"jwt_lang", claims.DeviceInfo.Language,
			"jwt_tz", claims.DeviceInfo.Timezone,
			"req_browser", secs.Browser,
			"req_os", secs.OperatingSystem,
			"req_lang", secs.Language,
			"req_tz", secs.Timezone)
		return &models.User{}, models.ErrSecurityMismatch
	}

	u.log.Info("Device validation successful", "email", claims.Email)

	return user, nil
}

func (u *controllerUser) validateToken(tkn *jwt.Token, err error) error {
	if errors.Is(err, jwt.ErrSignatureInvalid) {
		return u.logAndReturnError(models.ErrInvalidSignature.Error())
	}
	if errors.Is(err, jwt.ErrTokenExpired) {
		return u.logAndReturnError(models.ErrTokenExpired.Error())
	}
	if err != nil {
		return u.logAndReturnError(models.ErrParsingToken.Error())
	}
	if !tkn.Valid {
		return u.logAndReturnError(models.ErrInvalidToken.Error())
	}
	return nil
}

func (u *controllerUser) logAndReturnError(errMsg string) error {
	u.log.Error(errMsg)
	return errors.New(errMsg)
}

// validateDeviceInfo performs robust device fingerprinting validation
// allowing for minor differences that might occur due to browser updates or detection variations
func (u *controllerUser) validateDeviceInfo(jwtDevice, requestDevice models.DeviceInfo) bool {
	// Core security checks - these must match exactly
	if jwtDevice.Browser != requestDevice.Browser ||
		jwtDevice.OperatingSystem != requestDevice.OperatingSystem ||
		jwtDevice.Language != requestDevice.Language ||
		jwtDevice.Timezone != requestDevice.Timezone {
		u.log.Warn("Device validation failed on core fields",
			"jwt_browser", jwtDevice.Browser, "req_browser", requestDevice.Browser,
			"jwt_os", jwtDevice.OperatingSystem, "req_os", requestDevice.OperatingSystem,
			"jwt_lang", jwtDevice.Language, "req_lang", requestDevice.Language,
			"jwt_tz", jwtDevice.Timezone, "req_tz", requestDevice.Timezone)
		return false
	}

	// More flexible checks for version fields that might vary
	// Allow browser version to be similar (e.g., "119.0.0.0" vs "119.0.1.0")
	if !u.isVersionCompatible(jwtDevice.BrowserVersion, requestDevice.BrowserVersion) {
		u.log.Warn("Browser version incompatible",
			"jwt_version", jwtDevice.BrowserVersion, "req_version", requestDevice.BrowserVersion)
		return false
	}

	// Allow OS version to be similar
	if !u.isVersionCompatible(jwtDevice.OperatingSystemVersion, requestDevice.OperatingSystemVersion) {
		u.log.Warn("OS version incompatible",
			"jwt_version", jwtDevice.OperatingSystemVersion, "req_version", requestDevice.OperatingSystemVersion)
		return false
	}

	// CPU and cookies enabled should match, but be more lenient with CPU
	// CPU detection can vary between requests
	if jwtDevice.CookiesEnabled != requestDevice.CookiesEnabled {
		u.log.Warn("Cookies enabled mismatch",
			"jwt_cookies", jwtDevice.CookiesEnabled, "req_cookies", requestDevice.CookiesEnabled)
		return false
	}

	// CPU can be empty or vary, so only check if both are non-empty and different
	if jwtDevice.Cpu != "" && requestDevice.Cpu != "" && jwtDevice.Cpu != requestDevice.Cpu {
		u.log.Warn("CPU mismatch",
			"jwt_cpu", jwtDevice.Cpu, "req_cpu", requestDevice.Cpu)
		return false
	}

	return true
}

// isVersionCompatible checks if two version strings are compatible
// allowing for minor version differences
func (u *controllerUser) isVersionCompatible(version1, version2 string) bool {
	if version1 == version2 {
		return true
	}

	// For empty versions, consider them compatible
	if version1 == "" || version2 == "" {
		return true
	}

	// Simple version comparison - split by dots and compare major versions
	parts1 := strings.Split(version1, ".")
	parts2 := strings.Split(version2, ".")

	if len(parts1) == 0 || len(parts2) == 0 {
		return false
	}

	// Compare major version (first component)
	if parts1[0] != parts2[0] {
		return false
	}

	// If we have minor versions, compare them too (allow some flexibility)
	if len(parts1) > 1 && len(parts2) > 1 {
		minor1, err1 := strconv.Atoi(parts1[1])
		minor2, err2 := strconv.Atoi(parts2[1])

		// If we can parse minor versions, allow a difference of 1
		if err1 == nil && err2 == nil {
			diff := minor1 - minor2
			if diff < 0 {
				diff = -diff
			}
			return diff <= 1 // Allow minor version difference of 1
		}
	}

	return true
}
