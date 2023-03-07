package repository

import (
	"errors"

	"gorm.io/gorm"
)

// InjectDB :nodoc:
func (r *userRepository) InjectDB(db *gorm.DB) error {
	if db == nil {
		return errors.New("invalid db")
	}
	r.db = db
	return nil
}
