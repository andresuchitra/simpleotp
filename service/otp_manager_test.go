package service

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOTPManager(t *testing.T) {
	t.Setenv("OTP_SECRET", "mysecret")

	t.Run("Success length less than 4", func(t *testing.T) {
		// set env in testing env
		duration := int64(300000)
		mgr := NewOTPManager(0)

		assert.Equal(t, "mysecret", mgr.Secret)
		assert.Equal(t, DEFAULT_OTP_LENGTH, mgr.Length) // token less than 4 chars will be set to 4
		assert.Equal(t, duration, mgr.ExpiryDelay)      // delay will be 5 minutes = 300,000
	})

	t.Run("Success length more than 4", func(t *testing.T) {
		mgr := NewOTPManager(10)

		assert.Equal(t, "mysecret", mgr.Secret)
		assert.Equal(t, uint(10), mgr.Length)
	})
}

func TestOTPManagerGenerateOTP(t *testing.T) {
	t.Run("Success create OTP", func(t *testing.T) {
		myLength := uint(7)
		mgr := NewOTPManager(myLength)
		otp, err := mgr.GenerateOTP()

		assert.Nil(t, err)
		assert.Len(t, otp, int(myLength))
		assert.Regexp(t, regexp.MustCompile("^[0-9]"), otp) // otp will be digits only
	})
}
