package service

import (
	"context"
	"errors"
	"testing"

	"github.com/andresuchitra/simpleotp/models"
	mockRepository "github.com/andresuchitra/simpleotp/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateOTP(t *testing.T) {
	ctx := context.Background()

	t.Run("Success call CreateOTP", func(t *testing.T) {
		mockRepo := mockRepository.NewRepository(t)
		phone := "081233332222"

		mockRepo.On("CreateOTP", mock.Anything, mock.Anything).Return(nil).Once()

		service := NewOTPService(mockRepo)
		result, err := service.CreateOTP(&ctx, phone)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result, int(DEFAULT_OTP_LENGTH))
	})

	t.Run("Fail CreateOTP", func(t *testing.T) {
		mockRepo := mockRepository.NewRepository(t)
		phone := "081233332222"
		mockError := errors.New("DB error")

		mockRepo.On("CreateOTP", mock.Anything, mock.Anything).Return(mockError).Once()

		service := NewOTPService(mockRepo)
		result, err := service.CreateOTP(&ctx, phone)

		assert.NotNil(t, err)
		assert.EqualError(t, mockError, "DB error")
		assert.Equal(t, "", result) // if error, result from CreateOTP is ""
	})
}

func TestValidateOTP(t *testing.T) {
	ctx := context.Background()

	t.Run("Success ValidateOTP", func(t *testing.T) {
		mockRepo := mockRepository.NewRepository(t)
		phone := "+6281233332222"
		otp := "924495"

		mockOtpItem := models.OTPItem{
			Phone: phone,
			OTP:   otp,
		}

		mockRepo.On("FindByOTPTokenAndPhone", mock.Anything, otp, phone).Return(&mockOtpItem, nil).Once()
		mockRepo.On("UpdateOTPByID", mock.Anything, mock.AnythingOfType("string")).Return(nil).Once()

		service := NewOTPService(mockRepo)
		err := service.ValidateOTP(&ctx, otp, phone)

		assert.Nil(t, err)
	})

	t.Run("Error ValidateOTP Invalid token format", func(t *testing.T) {
		mockRepo := mockRepository.NewRepository(t)
		phone := "+6281233332222"
		otp := "333asddd33232" // non-digits otp is invalid

		service := NewOTPService(mockRepo)
		err := service.ValidateOTP(&ctx, otp, phone)

		assert.EqualError(t, err, "OTP format is invalid")
	})

	t.Run("Error ValidateOTP Invalid phone format", func(t *testing.T) {
		mockRepo := mockRepository.NewRepository(t)
		phone := "+62722222" // must be start with +628xxxx
		otp := "12222"

		service := NewOTPService(mockRepo)
		err := service.ValidateOTP(&ctx, otp, phone)

		assert.EqualError(t, err, "Phone format is invalid, must be +628xxxxx, between 11 to 12 digits")
	})
}
