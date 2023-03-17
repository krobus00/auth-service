package grpc

import (
	"errors"

	"github.com/krobus00/auth-service/internal/model"
)

// InjectUserUsecase :nodoc:
func (c *Server) InjectUserUsecase(usecase model.UserUsecase) error {
	if usecase == nil {
		return errors.New("invalid user usecase")
	}
	c.userUC = usecase
	return nil
}

func (c *Server) InjectAuthUsecase(usecase model.AuthUsecase) error {
	if usecase == nil {
		return errors.New("invalid auth usecase")
	}
	c.authUC = usecase
	return nil
}
