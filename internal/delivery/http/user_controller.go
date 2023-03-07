package http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/krobus00/auth-service/internal/constant"
	"github.com/krobus00/auth-service/internal/model"
	"github.com/labstack/echo/v4"
)

// UserController :nodoc:
type UserController struct {
	userUC model.UserUsecase
}

// NewUserController :nodoc:
func NewUserController() *UserController {
	return new(UserController)
}

// Register :nodoc:
func (c *UserController) Register(eCtx echo.Context) (err error) {
	var (
		ctx = eCtx.Request().Context()
		res = new(model.Response)
		req = new(model.HTTPUserRegistrationRequest)
	)

	err = eCtx.Bind(req)
	if err != nil {
		res = model.NewResponse().WithMessage("bad request")
		return eCtx.JSON(http.StatusBadRequest, res)
	}

	payload := req.ToPayload()

	token, err := c.userUC.Register(ctx, payload)
	if err != nil {
		val, ok := HTTPErrorMapping[err]
		if !ok {
			res = model.NewResponse().WithErrorMessage(ErrInternalServerError)
			return eCtx.JSON(http.StatusInternalServerError, res)
		}
		res = model.NewResponse().WithErrorMessage(val.err)
		return eCtx.JSON(val.statusCode, res)
	}

	res = model.NewResponse().WithData(token.ToHTTPResponse())
	return eCtx.JSON(http.StatusCreated, res)
}

// Login :nodoc:
func (c *UserController) Login(eCtx echo.Context) (err error) {
	var (
		ctx = eCtx.Request().Context()
		res = new(model.Response)
		req = new(model.HTTPUserLoginRequest)
	)

	err = eCtx.Bind(req)
	if err != nil {
		res = model.NewResponse().WithMessage("bad request")
		return eCtx.JSON(http.StatusBadRequest, res)
	}

	payload := req.ToPayload()

	token, err := c.userUC.Login(ctx, payload)
	if err != nil {
		val, ok := HTTPErrorMapping[err]
		if !ok {
			res = model.NewResponse().WithErrorMessage(ErrInternalServerError)
			return eCtx.JSON(http.StatusInternalServerError, res)
		}
		res = model.NewResponse().WithErrorMessage(val.err)
		return eCtx.JSON(val.statusCode, res)
	}

	res = model.NewResponse().WithData(token.ToHTTPResponse())
	return eCtx.JSON(http.StatusOK, res)
}

// GetUserInfo :nodoc:
func (c *UserController) GetUserInfo(eCtx echo.Context) (err error) {
	var (
		ctx = eCtx.Request().Context()
		res = new(model.Response)
	)

	user, err := c.userUC.GetUserInfo(ctx, &model.GetUserInfoPayload{
		ID: fmt.Sprintf("%v", eCtx.Get(string(constant.KeyUserIDCtx))),
	})
	if err != nil {
		val, ok := HTTPErrorMapping[err]
		if !ok {
			res = model.NewResponse().WithErrorMessage(ErrInternalServerError)
			return eCtx.JSON(http.StatusInternalServerError, res)
		}
		res = model.NewResponse().WithErrorMessage(val.err)
		return eCtx.JSON(val.statusCode, res)
	}
	res = model.NewResponse().WithData(user.ToHTTPResponse())
	return eCtx.JSON(http.StatusOK, res)
}

// RefreshToken :nodoc:
func (c *UserController) RefreshToken(eCtx echo.Context) (err error) {
	var (
		ctx = eCtx.Request().Context()
		res = new(model.Response)
	)

	tokenHeader := eCtx.Request().Header.Get("Authorization")
	tokenHeader = strings.Replace(tokenHeader, "Bearer ", "", -1)
	if tokenHeader == "" {
		return eCtx.JSON(http.StatusUnauthorized, res)
	}

	payload := &model.RefreshTokenPayload{
		UserID:  fmt.Sprintf("%v", eCtx.Get(string(constant.KeyUserIDCtx))),
		TokenID: fmt.Sprintf("%v", eCtx.Get(string(constant.KeyTokenIDCtx))),
	}

	token, err := c.userUC.RefreshToken(ctx, payload)
	if err != nil {
		val, ok := HTTPErrorMapping[err]
		if !ok {
			res = model.NewResponse().WithErrorMessage(ErrInternalServerError)
			return eCtx.JSON(http.StatusInternalServerError, res)
		}
		res = model.NewResponse().WithErrorMessage(val.err)
		return eCtx.JSON(val.statusCode, res)
	}

	res = model.NewResponse().WithData(token.ToHTTPResponse())
	return eCtx.JSON(http.StatusOK, res)
}

// Logout :nodoc:
func (c *UserController) Logout(eCtx echo.Context) (err error) {
	var (
		ctx = eCtx.Request().Context()
		res = new(model.Response)
	)

	tokenHeader := eCtx.Request().Header.Get("Authorization")
	tokenHeader = strings.Replace(tokenHeader, "Bearer ", "", -1)
	if tokenHeader == "" {
		return eCtx.JSON(http.StatusUnauthorized, res)
	}

	payload := &model.LogoutPayload{
		UserID:  fmt.Sprintf("%v", eCtx.Get(string(constant.KeyUserIDCtx))),
		TokenID: fmt.Sprintf("%v", eCtx.Get(string(constant.KeyTokenIDCtx))),
	}

	err = c.userUC.Logout(ctx, payload)
	if err != nil {
		val, ok := HTTPErrorMapping[err]
		if !ok {
			res = model.NewResponse().WithErrorMessage(ErrInternalServerError)
			return eCtx.JSON(http.StatusInternalServerError, res)
		}
		res = model.NewResponse().WithErrorMessage(val.err)
		return eCtx.JSON(val.statusCode, res)
	}

	res = model.NewDefaultResponse()
	return eCtx.JSON(http.StatusOK, res)
}
