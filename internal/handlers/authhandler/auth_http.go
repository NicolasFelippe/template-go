package authhandler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"template-go/internal/core/ports"
	"template-go/internal/handlers/helpers"
	"template-go/internal/handlers/usershandler"
)

type AuthHTTPHandler struct {
	sessionService ports.SessionService
}

func NewAuthHTTPHandler(sessionService ports.SessionService) *AuthHTTPHandler {
	return &AuthHTTPHandler{
		sessionService,
	}
}

func (hdl *AuthHTTPHandler) Login(ctx *gin.Context) {
	var req LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err))
		return
	}

	session, auth, err := hdl.sessionService.CreateSession(req.Username, req.Password, ctx.Request.UserAgent(), ctx.ClientIP())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse(err))
		return
	}
	var user usershandler.ResponseCreateUserDTO
	user.FromDomain(&session.User)
	rsp := &LoginUserResponse{
		User:                  user,
		ID:                    session.ID,
		AccessToken:           auth.AccessToken,
		RefreshToken:          session.RefreshToken,
		AccessTokenExpiresAt:  auth.AccessTokenExpiresAt,
		RefreshTokenExpiresAt: auth.RefreshTokenExpiresAt,
	}
	ctx.JSON(http.StatusOK, rsp)
}
