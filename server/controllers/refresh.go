package controllers

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/alvarotor/user-go/server/dto"
	"github.com/alvarotor/user-go/server/models"
	"github.com/golang-jwt/jwt/v5"
)

func (u *controllerUser) Refresh(ctx context.Context, refreshToken string) (int, *models.Token, error) {
	u.log.Info("🔍 [REFRESH_START] Refresh request initiated",
		"time", time.Now().UTC())

	claims := &dto.ClaimsRefreshResponse{}

	tkn, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (any, error) {
		return u.conf.JWTKey, nil
	})

	if err := u.validateToken(tkn, err); err != nil {
		u.log.Error("🔍 [REFRESH_ERROR] Refresh token validation failed",
			"error", err,
			"time", time.Now().UTC())
		return http.StatusBadRequest, &models.Token{}, err
	}

	u.log.Info("🔍 [REFRESH_TOKEN_VALID] Refresh token parsed successfully",
		"expires_at", claims.ExpiresAt,
		"time_remaining", time.Until(claims.ExpiresAt.Time),
		"time", time.Now().UTC())

	user, err := u.GetByCodeRefresh(ctx, claims.CodeRefresh)
	if errors.Is(err, models.ErrUserNotFound) {
		return http.StatusNotFound, &models.Token{}, models.ErrInvalidCode
	}
	if err != nil {
		return http.StatusInternalServerError, &models.Token{}, err
	}

	if user == nil {
		errMsg := "code refresh is invalid"
		u.log.Error(errMsg)
		return http.StatusBadRequest, &models.Token{}, errors.New(errMsg)
	}

	if user.Code == "OUT" || strings.TrimSpace(user.Code) == "" {
		errMsg := "code refresh is invalid"
		u.log.Error(errMsg)
		return http.StatusBadRequest, &models.Token{}, errors.New(errMsg)
	}

	if u.conf.SizeRandomStringValidationRefresh != len(user.CodeRefresh) {
		errMsg := "code refresh is invalid"
		u.log.Error(errMsg)
		return http.StatusBadRequest, &models.Token{}, errors.New(errMsg)
	}

	// Generate new refresh code and update atomically to prevent race conditions
	newCodeRefresh := u.GenerateRandomString(u.conf.SizeRandomStringValidationRefresh)

	// Update the code_refresh - this should be atomic at the database level
	// If multiple requests try to update the same user simultaneously, only one will succeed
	u.log.Info("🔍 [REFRESH_UPDATE_START] Updating refresh code",
		"user_id", user.ID,
		"time", time.Now().UTC())
	err = u.UpdateField(ctx, user.ID, "code_refresh", newCodeRefresh)
	if err != nil {
		u.log.Error("🔍 [REFRESH_UPDATE_ERROR] Failed to update refresh code",
			"error", err,
			"user_id", user.ID,
			"time", time.Now().UTC())
		return http.StatusInternalServerError, &models.Token{}, err
	}

	u.log.Info("🔍 [REFRESH_UPDATE_SUCCESS] Refresh code updated",
		"user_id", user.ID,
		"old_code", user.CodeRefresh[:8]+"...",
		"new_code", newCodeRefresh[:8]+"...",
		"time", time.Now().UTC())

	// Update user object with new code for token generation
	user.CodeRefresh = newCodeRefresh

	u.log.Info("🔍 [REFRESH_VALIDATE_CALL] Calling Validate to generate new tokens",
		"time", time.Now().UTC())
	status, modelToken, err := u.Validate(ctx, user.Code)
	if err != nil {
		u.log.Error("🔍 [REFRESH_VALIDATE_ERROR] Validate call failed",
			"error", err,
			"time", time.Now().UTC())
		return http.StatusBadRequest, &models.Token{}, err
	}

	u.log.Info("🔍 [REFRESH_SUCCESS] Refresh completed successfully",
		"status", status,
		"new_access_expires", modelToken.TokenExpires,
		"new_refresh_expires", modelToken.TokenRefreshExpires,
		"time", time.Now().UTC())

	return status, &modelToken, nil
}
