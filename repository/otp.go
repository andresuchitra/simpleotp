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
	FindByOTPTokenAndPhone(ctx *context.Context, token string, phone string) (*models.OTPItem, error)
	UpdateOTPByID(ctx *context.Context, id string) error
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

func (r repo) FindByOTPTokenAndPhone(ctx *context.Context, token string, phone string) (*models.OTPItem, error) {
	var item models.OTPItem

	result := r.DB.Where("otp = ? and phone = ?", token, phone).First(&item)
	if result.Error != nil {
		return nil, result.Error
	}
	return &item, nil
}

func (r repo) UpdateOTPByID(ctx *context.Context, id string) error {
	result := r.DB.Table("otp_items").Where("id = ?", id).Update("is_used", true)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
