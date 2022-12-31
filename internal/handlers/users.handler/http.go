package users_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
	"template-go/internal/core/ports"
)

type UserHTTPHandler struct {
	userService ports.UserService
}

func NewUserHTTPHandler(userService ports.UserService) *UserHTTPHandler {
	return &UserHTTPHandler{
		userService: userService,
	}
}

func (hdl *UserHTTPHandler) CreateUser(ctx *gin.Context) {
	var req RequestUserDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := hdl.userService.CreateUser(
		req.Username,
		req.Password,
		req.FullName,
		req.Email,
	)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var rsp ResponseCreateUserDTO
	rsp.FromDomain(result)

	ctx.JSON(http.StatusOK, rsp)
}

func (hdl *UserHTTPHandler) ListUsers(ctx *gin.Context) {
	var req listUsersRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	offset := (req.PageID - 1) * req.PageSize
	result, err := hdl.userService.ListUsers(
		&req.PageSize,
		&offset,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	//var rsp ResponseCreateUserDTO
	//rsp.FromDomain(result)

	ctx.JSON(http.StatusOK, result)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
