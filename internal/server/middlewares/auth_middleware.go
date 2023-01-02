package middlewares

import (
	"errors"
	"fmt"
	"strings"
	domainErrors "template-go/internal/core/domain/errors"
	"template-go/pkg/makertoken"

	"github.com/gin-gonic/gin"
)

const (
	AuthorizationHeaderKey  = "authorization"
	AuthorizationTypeBearer = "bearer"
	AuthorizationPayloadKey = "authorization_payload"
)

func AuthMiddleware(tokenMaker makertoken.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(AuthorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			appError := domainErrors.NewAppError(err, domainErrors.NotAuthenticated)
			_ = ctx.Error(appError)
			ctx.Abort()
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authroization header format")
			appError := domainErrors.NewAppError(err, domainErrors.NotAuthenticated)
			_ = ctx.Error(appError)
			ctx.Abort()
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != AuthorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			appError := domainErrors.NewAppError(err, domainErrors.NotAuthenticated)
			_ = ctx.Error(appError)
			ctx.Abort()
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			appError := domainErrors.NewAppError(err, domainErrors.NotAuthenticated)
			_ = ctx.Error(appError)
			ctx.Abort()
			return
		}

		ctx.Set(AuthorizationPayloadKey, payload)
	}
}
