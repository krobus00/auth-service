package usecase

import (
	"errors"

	"github.com/krobus00/auth-service/internal/model"
	"gorm.io/gorm"
)

// InjectUserRepo :nodoc:
func (uc *userUsecase) InjectUserRepo(repo model.UserRepository) error {
	if repo == nil {
		return errors.New("invalid user repo")
	}
	uc.userRepo = repo
	return nil
}

// InjectTokenRepo :nodoc:
func (uc *userUsecase) InjectTokenRepo(repo model.TokenRepository) error {
	if repo == nil {
		return errors.New("invalid token repo")
	}
	uc.tokenRepo = repo
	return nil
}

// InjectDB :nodoc:
func (uc *userUsecase) InjectDB(db *gorm.DB) error {
	if db == nil {
		return errors.New("invalid db")
	}
	uc.db = db
	return nil
}
