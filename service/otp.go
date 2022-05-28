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

const (
	PHONE_REGEX_INDONESIA = `^\+628[0-9]{9,10}$`

	ERR_OTP_ALREADY_USED       = "OTP has been used. Please create again"
	ERR_INVALID_PHONE_FORMAT   = "Phone format is invalid, must be +628xxxxx, between 11 to 12 digits"
	ERR_INVALID_PAYLOAD_FORMAT = "OTP format is invalid"
	ERR_DIFFERENT_PHONE_NUMBER = "Something wrong on phone data. Please create again"
	ERR_EXPIRED_OTP            = "OTP has been expired. Please create again"
)

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

func (s service) CreateOTP(ctx *context.Context, phone string) (string, error) {
	var newOtp models.OTPItem

	// validate phone
	// validate phone
	regx, err := regexp.Compile(PHONE_REGEX_INDONESIA)
	if err != nil {
		return "", errors.New("Error checking phone format")
	}

	if !regx.MatchString(phone) {
		return "", errors.New(ERR_INVALID_PHONE_FORMAT)
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

func (s service) ValidateOTP(ctx *context.Context, otpToken string, phone string) error {
	// validate token. must be all digits
	_, err := strconv.Atoi(otpToken)
	if err != nil {
		return errors.New(ERR_INVALID_PAYLOAD_FORMAT)
	}

	// validate phone
	regx, err := regexp.Compile(PHONE_REGEX_INDONESIA)
	if err != nil {
		return errors.New("Error checking phone format")
	}

	if !regx.MatchString(phone) {
		return errors.New(ERR_INVALID_PHONE_FORMAT)
	}

	// query to DB
	newOtp, err := s.repo.FindByOTPTokenAndPhone(ctx, otpToken, phone)
	if err != nil {
		return err
	}

	// validate if is_used is false, raise error if it's already used
	if newOtp.IsUsed {
		return errors.New(ERR_OTP_ALREADY_USED)
	}

	// validate if phone in created record is different to request
	// raise error if it's different
	if newOtp.Phone != phone {
		return errors.New(ERR_DIFFERENT_PHONE_NUMBER)
	}

	// if expiry time is passed, raise error
	if newOtp.ExpiryAt < time.Now().UnixMilli() {
		return errors.New(ERR_EXPIRED_OTP)
	}

	// update to DB
	err = s.repo.UpdateOTPByID(ctx, newOtp.ID)
	if err != nil {
		return err
	}

	return nil
}
