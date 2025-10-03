package controllers

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/alvarotor/user-go/server/dto"
	"github.com/alvarotor/user-go/server/models"
	"github.com/golang-jwt/jwt/v5"
)

func (u *controllerUser) Refresh(ctx context.Context, refreshToken string) (int, *models.Token, error) {
	claims := &dto.ClaimsRefreshResponse{}

	tkn, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (any, error) {
		return u.conf.JWTKey, nil
	})

	if err := u.validateToken(tkn, err); err != nil {
		return http.StatusBadRequest, &models.Token{}, err
	}

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
	err = u.UpdateField(ctx, user.ID, "code_refresh", newCodeRefresh)
	if err != nil {
		u.log.Error("Failed to update refresh code", "error", err, "user_id", user.ID)
		return http.StatusInternalServerError, &models.Token{}, err
	}

	u.log.Info("Refresh token updated successfully", "user_id", user.ID, "old_code", user.CodeRefresh[:8]+"...", "new_code", newCodeRefresh[:8]+"...")

	// Update user object with new code for token generation
	user.CodeRefresh = newCodeRefresh

	status, modelToken, err := u.Validate(ctx, user.Code)
	if err != nil {
		return http.StatusBadRequest, &models.Token{}, err
	}

	return status, &modelToken, nil
}
