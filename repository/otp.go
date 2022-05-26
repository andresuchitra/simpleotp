package repository

import (
	"context"

	"github.com/andresuchitra/simpleotp/models"
	"gorm.io/gorm"
)

type OTPRepository struct {
	DB *gorm.DB
}

type Repository interface {
	CreateOTP(ctx *context.Context, newOtp *models.OTPItem) error
}

func NewRepository(db *gorm.DB) *OTPRepository {
	db.AutoMigrate(&models.OTPItem{})

	return &OTPRepository{
		DB: db,
	}
}

func (r *OTPRepository) CreateOTP(ctx *context.Context, newOtp *models.OTPItem) error {
	result := r.DB.Create(newOtp)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *OTPRepository) FindByOTPToken(ctx *context.Context, token string) error {
	// result := r.DB.Find(&)
	// if result.Error != nil {
	// 	return result.Error
	// }
	return nil
}
