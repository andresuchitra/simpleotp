package repository

import (
	"context"

	"github.com/andresuchitra/simpleotp/models"
	"gorm.io/gorm"
)

type repo struct {
	DB *gorm.DB
}

type Repository interface {
	CreateOTP(ctx *context.Context, newOtp *models.OTPItem) error
}

func NewRepository(db *gorm.DB) Repository {
	db.AutoMigrate(&models.OTPItem{})

	return repo{
		DB: db,
	}
}

func (r repo) CreateOTP(ctx *context.Context, newOtp *models.OTPItem) error {
	result := r.DB.Create(newOtp)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r repo) FindByOTPToken(ctx *context.Context, token string) error {
	// result := r.DB.Find(&)
	// if result.Error != nil {
	// 	return result.Error
	// }
	return nil
}
