package service

import (
	"crypto/rand"
	"fmt"
	"os"
	"time"
)

const (
	digits             = "1234567890"
	DEFAULT_OTP_LENGTH = uint(6)
)

type OTPManager struct {
	Length      uint
	Secret      string
	ExpiryDelay int64
}

func NewOTPManager(length uint) *OTPManager {
	if length < 4 {
		length = DEFAULT_OTP_LENGTH // valid otp length only 4 or more
	}

	secret := os.Getenv("OTP_SECRET")
	delayExpiry := 5 * time.Minute

	return &OTPManager{
		Length:      length,
		Secret:      secret,
		ExpiryDelay: delayExpiry.Milliseconds(),
	}
}

func (g *OTPManager) GenerateOTP() (string, error) {
	otpCharsLength := len(digits)
	buffer := make([]byte, g.Length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", fmt.Errorf("Error reading buffer: %v", err)
	}

	for i := 0; i < int(g.Length); i++ {
		buffer[i] = digits[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}
