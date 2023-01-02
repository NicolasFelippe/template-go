package authhandler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	domainErrors "template-go/internal/core/domain/errors"
	"template-go/internal/core/domain/sessions"
)

type AuthHTTPHandler struct {
	sessionService sessions.SessionService
}

func NewAuthHTTPHandler(sessionService sessions.SessionService) *AuthHTTPHandler {
	return &AuthHTTPHandler{
		sessionService,
	}
}

func (hdl *AuthHTTPHandler) Login(ctx *gin.Context) {
	var req LoginValidator
	if err := ctx.ShouldBindJSON(&req); err != nil {
		appError := domainErrors.NewAppError(err, domainErrors.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	session, auth, err := hdl.sessionService.CreateSession(
		req.Username,
		req.Password,
		ctx.Request.UserAgent(),
		ctx.ClientIP(),
	)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	rsp := toResponseModel(auth, session)
	ctx.JSON(http.StatusOK, rsp)
}
