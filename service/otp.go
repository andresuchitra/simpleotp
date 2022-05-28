package service

import (
	"context"
	"errors"
	"regexp"
	"strconv"
	"time"

	"github.com/andresuchitra/simpleotp/models"
	"github.com/andresuchitra/simpleotp/repository"
	"github.com/google/uuid"
)

const PHONE_REGEX_INDONESIA = `^\+628[0-9]{9,10}$`

type service struct {
	repo    repository.Repository
	manager OTPManager
}

func NewOTPService(repo repository.Repository) service {
	// generate manager
	otpManager := NewOTPManager(6)

	newService := service{
		repo:    repo,
		manager: *otpManager,
	}

	return newService
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

func (s *service) ValidateOTP(ctx *context.Context, otpToken string, phone string) error {
	// validate token. must be all digits
	_, err := strconv.Atoi(otpToken)
	if err != nil {
		return errors.New("OTP format is invalid")
	}

	// validate phone
	regx, err := regexp.Compile(PHONE_REGEX_INDONESIA)
	if err != nil {
		return errors.New("Error checking phone format")
	}

	if !regx.MatchString(phone) {
		return errors.New("Phone format is invalid, must be +628xxxxx, between 11 to 12 digits")
	}

	// query to DB
	newOtp, err := s.repo.FindByOTPTokenAndPhone(ctx, otpToken, phone)
	if err != nil {
		return err
	}

	// validate if is_used is false, raise error if it's already used
	if newOtp.IsUsed {
		return errors.New("OTP has been used. Please create again")
	}

	// if expiry time is passed, raise error
	if newOtp.ExpiryAt > time.Now().UnixMilli() {
		return errors.New("OTP has been expired. Please create again")
	}

	// update to DB
	err = s.repo.UpdateOTPByID(ctx, newOtp.ID)
	if err != nil {
		return err
	}

	return nil
}
