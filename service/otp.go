package service

import (
	"context"
	"errors"
	"time"

	"github.com/andresuchitra/simpleotp/models"
	"github.com/andresuchitra/simpleotp/repository"
	"github.com/google/uuid"
)

type service struct {
	repo    repository.OTPRepository
	manager OTPManager
}

func NewOTPService(repo *repository.OTPRepository) OTPService {
	// generate manager
	otpManager := NewOTPManager(6)

	newService := service{
		repo:    *repo,
		manager: *otpManager,
	}

	return &newService
}

func (s *service) CreateOTP(ctx *context.Context, phone string) (string, error) {
	var newOtp models.OTPItem

	// validate phone
	if phone == "" {
		return "", errors.New("invalid phone data")
	}
	newOtp.Phone = phone

	// generate random UUID as index in DB
	uuid := uuid.New()
	newOtp.ID = uuid.String()

	otpToken, err := s.manager.GenerateOTP()
	if err != nil {
		return "", err
	}

	// store the otp to current mobile phone request
	newOtp.ExpiryAt = time.Now().UnixMilli() + s.manager.ExpiryDelay
	// set otpToken
	newOtp.OTP = otpToken

	// store to DB
	err = s.repo.CreateOTP(ctx, &newOtp)
	if err != nil {
		return "", err
	}

	return otpToken, nil
}

func (s *service) ValidateOTP(ctx *context.Context, otpToken string) error {
	// var newOtp models.OTPItem

	// validate token first. must be integer
	// _, err := strconv.Atoi(otpToken)
	// if err != nil {
	// 	return errors.New("OTP format is invalid")
	// }

	// // query to DB
	// err = s.repo.FindByOTPToken(ctx, otpToken)
	// if err != nil {
	// 	return err
	// }

	return nil
}
