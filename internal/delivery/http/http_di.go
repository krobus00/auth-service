package http

import (
	"errors"

	"github.com/labstack/echo/v4"
)

// InjectEcho :nodoc:
func (d *HTTPDelivery) InjectEcho(e *echo.Echo) error {
	if e == nil {
		return errors.New("invalid echo")
	}
	d.e = e
	return nil
}

// InjectHTTPMiddleware :nodoc:
func (d *HTTPDelivery) InjectHTTPMiddleware(m *HTTPMiddleware) error {
	if m == nil {
		return errors.New("invalid http middlewar")
	}
	d.middleware = m
	return nil
}

// InjectUserController :nodoc:
func (d *HTTPDelivery) InjectUserController(c *UserController) error {
	if c == nil {
		return errors.New("invalid user controller")
	}
	d.userController = c
	return nil
}
