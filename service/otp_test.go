package service

import (
	"context"
	"testing"

	mockRepository "github.com/andresuchitra/simpleotp/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateOTP(t *testing.T) {
	ctx := context.Background()
	mockRepo := mockRepository.NewRepository(t)

	t.Run("Success call CreateOTP", func(t *testing.T) {
		phone := "081233332222"

		mockRepo.On("CreateOTP", mock.Anything, mock.Anything).Return(nil).Once()

		service := NewOTPService(mockRepo)
		result, err := service.CreateOTP(&ctx, phone)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Len(t, result, int(DEFAULT_OTP_LENGTH))
	})
}
