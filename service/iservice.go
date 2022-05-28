package service

import "context"

type OTPService interface {
	CreateOTP(ctx *context.Context, phone string) (string, error)
	ValidateOTP(ctx *context.Context, otpToken string, phone string) error
}
