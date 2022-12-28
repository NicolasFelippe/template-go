package usershandler

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
	//hashedPassword, err := util.HashPassword(req.Password)
	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	//	return
	//}
	rsp, err := hdl.userService.CreateUser(
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

	ctx.JSON(http.StatusOK, rsp)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
