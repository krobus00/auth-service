package repository

import (
	"context"
	"errors"

	"github.com/krobus00/auth-service/internal/model"
	"github.com/krobus00/auth-service/internal/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository :nodoc:
func NewUserRepository() model.UserRepository {
	return new(userRepository)
}

// Create :nodoc:
func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	logger := log.WithFields(log.Fields{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
	})

	db := utils.GetTxFromContext(ctx, r.db)

	err := db.WithContext(ctx).Create(user).Error
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return nil
}

// FindByID :nodoc:
func (r *userRepository) FindByID(ctx context.Context, id string) (*model.User, error) {
	logger := log.WithFields(log.Fields{
		"id": id,
	})
	user := new(model.User)

	db := utils.GetTxFromContext(ctx, r.db)
	err := db.WithContext(ctx).Take(user, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, err
	}

	return user, nil
}

// FindByUsername :nodoc:
func (r *userRepository) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	logger := log.WithFields(log.Fields{
		"username": username,
	})
	user := new(model.User)

	db := utils.GetTxFromContext(ctx, r.db)
	err := db.WithContext(ctx).Take(user, "username = ?", username).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, err
	}

	return user, nil
}

// FindByEmail:nodoc:
func (r *userRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	logger := log.WithFields(log.Fields{
		"email": email,
	})
	user := new(model.User)

	db := utils.GetTxFromContext(ctx, r.db)
	err := db.WithContext(ctx).Take(user, "email = ?", email).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		logger.Error(err.Error())
		return nil, err
	}

	return user, nil
}

func (r *userRepository) UpdateByID(ctx context.Context, id string) (*model.User, error) {
	return nil, errors.New("unimplemented")
}

func (r *userRepository) DeleteByID(ctx context.Context, id string) error {
	return errors.New("unimplemented")
}
