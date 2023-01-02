package handlers

import (
	"net/http"
	"strings"
	domain "template-go/internal/core/domain/errors"

	"github.com/gin-gonic/gin"
)

// Handler is Gin middleware to handle errors.
func Handler(c *gin.Context) {
	// Execute request handlers and then handle any errors
	c.Next()
	errs := c.Errors

	if len(errs) > 0 {
		appError, ok := errs[0].Err.(*domain.AppError)
		if ok {
			switch strings.ToLower(appError.Type) {
			case strings.ToLower(domain.NotFound):
				c.AbortWithStatusJSON(http.StatusNotFound, appError.Error())
				return
			case strings.ToLower(domain.ValidationError):
				c.AbortWithStatusJSON(http.StatusBadRequest, appError.Error())
				return
			case strings.ToLower(domain.ResourceAlreadyExists):
				c.AbortWithStatusJSON(http.StatusConflict, appError.Error())
				return
			case strings.ToLower(domain.NotAuthenticated):
				c.AbortWithStatusJSON(http.StatusUnauthorized, appError.Error())
				return
			case strings.ToLower(domain.NotAuthorized):
				c.AbortWithStatusJSON(http.StatusForbidden, appError.Error())
				return
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, appError.Error())
				return
			}
		}

		// Error is not AppError, return a generic internal server error
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Internal Server Errror")
		return
	}
}
