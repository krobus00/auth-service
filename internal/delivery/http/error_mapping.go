package http

import (
	"errors"
	"net/http"

	"github.com/krobus00/auth-service/internal/model"
)

type httpError struct {
	err        error
	statusCode int
}

var ErrInternalServerError = errors.New("internal server error")

var HTTPErrorMapping = map[error]httpError{
	model.ErrUserNotFound: {
		err:        model.ErrUserNotFound,
		statusCode: http.StatusNotFound,
	},
	model.ErrInvalidTokenType: {
		err:        model.ErrInvalidTokenType,
		statusCode: http.StatusBadRequest,
	},
	model.ErrTokenInvalid: {
		err:        model.ErrTokenInvalid,
		statusCode: http.StatusUnauthorized,
	},
	model.ErrUsernameOrEmailAlreadyTaken: {
		err:        model.ErrUsernameOrEmailAlreadyTaken,
		statusCode: http.StatusBadRequest,
	},
	model.ErrWrongUsernameOrPassword: {
		err:        model.ErrWrongUsernameOrPassword,
		statusCode: http.StatusBadRequest,
	},
}
