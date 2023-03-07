package http

import (
	"errors"

	"github.com/krobus00/auth-service/internal/model"
)

// InjectUserUsecase :nodoc:
func (c *UserController) InjectUserUsecase(usecase model.UserUsecase) error {
	if usecase == nil {
		return errors.New("invalid user usecase")
	}
	c.userUC = usecase
	return nil
}
