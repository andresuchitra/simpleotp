package repository

import (
	"context"

	"github.com/andresuchitra/simpleotp/models"
	"gorm.io/gorm"
)

type OTPRepository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *OTPRepository {
	return &OTPRepository{
		DB: db,
	}
}

func (r *OTPRepository) CreateOTP(ctx context.Context, newOtp *models.OTPItem) error {
	result := r.DB.Create(newOtp)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
