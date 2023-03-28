package grpc

import (
	"errors"

	"github.com/krobus00/auth-service/internal/model"
)

func (t *Server) InjectUserUsecase(usecase model.UserUsecase) error {
	if usecase == nil {
		return errors.New("invalid user usecase")
	}
	t.userUC = usecase
	return nil
}

func (t *Server) InjectAuthUsecase(usecase model.AuthUsecase) error {
	if usecase == nil {
		return errors.New("invalid auth usecase")
	}
	t.authUC = usecase
	return nil
}
