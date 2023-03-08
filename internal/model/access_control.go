package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

const (
	FullAccess = "FULL_ACCESS"
)

type AccessControl struct {
	ID        string
	Name      string
	HasAccess bool `gorm:"->"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (AccessControl) TableName() string {
	return "access_controls"
}

type AccessControlRepository interface {
	Create(ctx context.Context, ac *AccessControl) error
	Remove(ctx context.Context, id string) error
	UpdateByID(ctx context.Context, ac *AccessControl) error

	// DI
	InjectDB(db *gorm.DB) error
}
