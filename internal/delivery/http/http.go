package http

import (
	"github.com/krobus00/auth-service/internal/model"
	"github.com/labstack/echo/v4"
)

// HTTPDelivery :nodoc:
type HTTPDelivery struct {
	e              *echo.Echo
	middleware     *HTTPMiddleware
	userController *UserController
}

// NewHTTPDelivery :nodoc:
func NewHTTPDelivery() *HTTPDelivery {
	return new(HTTPDelivery)
}

func (d *HTTPDelivery) InitRoutes() {
	api := d.e.Group("/api")

	users := api.Group("/users")
	users.POST("/register", d.userController.Register)
	users.POST("/login", d.userController.Login)
	users.GET("/me", d.userController.GetUserInfo, d.middleware.DecodeJWTToken(model.AccessToken))

	auth := api.Group("/auth")
	auth.GET("/refresh-token", d.userController.RefreshToken, d.middleware.DecodeJWTToken(model.RefreshToken))
	auth.DELETE("/logout", d.userController.Logout, d.middleware.DecodeJWTToken(model.AccessToken))
}
