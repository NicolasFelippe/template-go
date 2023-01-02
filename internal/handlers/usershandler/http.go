package usershandler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	domainErrors "template-go/internal/core/domain/errors"
	"template-go/internal/core/domain/users"
)

type UserHTTPHandler struct {
	userService users.UserService
}

func NewUserHTTPHandler(userService users.UserService) *UserHTTPHandler {
	return &UserHTTPHandler{
		userService,
	}
}

func (hdl *UserHTTPHandler) CreateUser(ctx *gin.Context) {
	user, err := BindJson(ctx)
	if err != nil {
		appError := domainErrors.NewAppError(err, domainErrors.ValidationError)
		_ = ctx.Error(appError)
		return
	}
	result, err := hdl.userService.CreateUser(user)

	if err != nil {
		_ = ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, *toResponseModel(result))
}

func (hdl *UserHTTPHandler) ListUsers(ctx *gin.Context) {
	var req PageValidator
	if err := ctx.ShouldBindQuery(&req); err != nil {
		appError := domainErrors.NewAppError(err, domainErrors.ValidationError)
		_ = ctx.Error(appError)
		return
	}

	//authPayload := ctx.MustGet(middlewares.AuthorizationPayloadKey).(*makertoken.Payload)
	//logger.Logger.Info(authPayload.Username)
	offset := (req.PageID - 1) * req.PageSize
	results, err := hdl.userService.ListUsersByPagination(
		&req.PageSize,
		&offset,
	)

	if err != nil {
		_ = ctx.Error(err)
		return
	}

	var responseItems = make([]UserResponse, len(results))

	for i, element := range results {
		responseItems[i] = *toResponseModel(&element)
	}
	//TODO: adicionar atributos ao response referente a paginação caso necessario.
	ctx.JSON(http.StatusOK, responseItems)
}
