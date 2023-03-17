package usecase

import (
	"context"
	"errors"

	"github.com/krobus00/auth-service/internal/model"
	"github.com/krobus00/auth-service/internal/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type userUsecase struct {
	userRepo  model.UserRepository
	tokenRepo model.TokenRepository
	db        *gorm.DB
}

// NewUserUsecase :nodoc:
func NewUserUsecase() model.UserUsecase {
	return new(userUsecase)
}

// Register :nodoc:
func (uc *userUsecase) Register(ctx context.Context, payload *model.UserRegistrationPayload) (*model.AuthResponse, error) {
	logger := log.WithFields(log.Fields{
		"username": payload.Username,
		"email":    payload.Email,
	})

	isUsernameOrEmailExist, err := uc.isUsernameOrEmailExist(ctx, payload.Username, payload.Email)
	if err != nil {
		return nil, err
	}
	if isUsernameOrEmailExist {
		return nil, model.ErrUsernameOrEmailAlreadyTaken
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	tx := uc.db.Begin()
	ctx = utils.NewTxContext(ctx, tx)

	newUser := &model.User{
		ID:       utils.GenerateUUID(),
		Username: payload.Username,
		Email:    payload.Email,
		Password: hashedPassword,
	}

	err = uc.userRepo.Create(ctx, newUser)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	token, err := uc.generateToken(ctx, newUser.ID)
	if err != nil {
		return nil, err
	}

	return token, tx.Commit().Error
}

// Login :nodoc:
func (uc *userUsecase) Login(ctx context.Context, payload *model.UserLoginPayload) (*model.AuthResponse, error) {
	logger := log.WithFields(log.Fields{
		"username": payload.Username,
	})

	user, err := uc.findUserByUsernameOrEmail(ctx, payload.Username)
	if err != nil {
		logger.Error(err.Error())
		if errors.Is(err, model.ErrUserNotFound) {
			return nil, model.ErrWrongUsernameOrPassword
		}
		return nil, err
	}

	err = utils.ComparePassword(user.Password, payload.Password)
	if err != nil {
		return nil, model.ErrWrongUsernameOrPassword
	}

	token, err := uc.generateToken(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return token, nil
}

// GetUserInfo :nodoc:
func (uc *userUsecase) GetUserInfo(ctx context.Context, payload *model.GetUserInfoPayload) (*model.UserInfoResponse, error) {
	logger := log.WithFields(log.Fields{
		"id": payload.ID,
	})
	user, err := uc.userRepo.FindByID(ctx, payload.ID)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	if user == nil {
		return nil, model.ErrUserNotFound
	}
	return &model.UserInfoResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}, nil
}

// RefreshToken :nodoc:
func (uc *userUsecase) RefreshToken(ctx context.Context, payload *model.RefreshTokenPayload) (*model.AuthResponse, error) {
	logger := log.WithFields(log.Fields{
		"userID": payload.UserID,
	})

	isValidToken, err := uc.tokenRepo.IsValidToken(ctx, payload.UserID, payload.TokenID, model.RefreshToken)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	if !isValidToken {
		return nil, model.ErrTokenInvalid
	}

	_ = uc.tokenRepo.Revoke(ctx, payload.UserID, payload.TokenID, model.AccessToken)
	_ = uc.tokenRepo.Revoke(ctx, payload.UserID, payload.TokenID, model.RefreshToken)

	token, err := uc.generateToken(ctx, payload.UserID)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return token, nil
}

// Logout :nodoc:
func (uc *userUsecase) Logout(ctx context.Context, payload *model.LogoutPayload) error {
	_ = uc.tokenRepo.Revoke(ctx, payload.UserID, payload.TokenID, model.AccessToken)
	_ = uc.tokenRepo.Revoke(ctx, payload.UserID, payload.TokenID, model.RefreshToken)

	return nil
}

func (uc *userUsecase) generateToken(ctx context.Context, userID string) (*model.AuthResponse, error) {
	tokenID := utils.GenerateUUID()
	accessToken, err := uc.tokenRepo.Create(ctx, userID, tokenID, model.AccessToken)
	if err != nil {
		return nil, err
	}
	refreshToken, err := uc.tokenRepo.Create(ctx, userID, tokenID, model.RefreshToken)
	if err != nil {
		return nil, err
	}
	return &model.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (uc *userUsecase) findUserByUsernameOrEmail(ctx context.Context, username string) (*model.User, error) {
	user, err := uc.userRepo.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return user, nil
	}
	user, err = uc.userRepo.FindByEmail(ctx, username)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return user, nil
	}
	return nil, model.ErrUserNotFound
}

func (uc *userUsecase) isUsernameOrEmailExist(ctx context.Context, username string, email string) (bool, error) {
	user, err := uc.userRepo.FindByUsername(ctx, username)
	if err != nil {
		return false, err
	}
	if user != nil {
		return true, nil
	}
	user, err = uc.userRepo.FindByEmail(ctx, username)
	if err != nil {
		return false, err
	}
	if user != nil {
		return true, nil
	}
	return false, nil
}
