package controllers

import (
	"context"
	"net/http"
)

func (u *controllerUser) Health(c context.Context, code uint32) int {
	if code != 1 {
		return http.StatusInternalServerError
	}
	return http.StatusOK
}
