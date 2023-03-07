package http

import (
	"net/http"
	"strings"

	"github.com/krobus00/auth-service/internal/constant"
	"github.com/krobus00/auth-service/internal/model"
	"github.com/krobus00/auth-service/internal/utils"
	"github.com/labstack/echo/v4"
)

// HTTPMiddleware :nodoc:
type HTTPMiddleware struct {
	tokenRepo model.TokenRepository
}

// NewHTTPMiddleware :nodoc:
func NewHTTPMiddleware() *HTTPMiddleware {
	return new(HTTPMiddleware)
}

// DecodeJWTToken :nodoc:
func (m *HTTPMiddleware) DecodeJWTToken(tokenType model.TokenType) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(eCtx echo.Context) error {
			ctx := eCtx.Request().Context()

			res := model.NewResponse().WithMessage(model.ErrTokenInvalid.Error())
			tokenHeader := eCtx.Request().Header.Get("Authorization")
			tokenHeader = strings.Replace(tokenHeader, "Bearer ", "", -1)
			if tokenHeader == "" {
				return eCtx.JSON(http.StatusUnauthorized, res)
			}

			token, err := utils.ParseToken(tokenHeader)
			if err != nil {
				return eCtx.JSON(http.StatusUnauthorized, res)
			}
			userID, err := utils.GetUserID(token)
			if err != nil {
				return eCtx.JSON(http.StatusUnauthorized, res)
			}

			tokenID, err := utils.GetTokenID(token)
			if err != nil {
				return eCtx.JSON(http.StatusUnauthorized, res)
			}

			isValid, err := m.tokenRepo.IsValidToken(ctx, userID, tokenID, tokenType)
			if err != nil || !isValid {
				return eCtx.JSON(http.StatusUnauthorized, res)
			}

			eCtx.Set(string(constant.KeyUserIDCtx), userID)
			eCtx.Set(string(constant.KeyTokenIDCtx), tokenID)
			return next(eCtx)
		}
	}
}
