package service

import "context"

type OTPService interface {
	CreateOTP(ctx *context.Context, phone string) (string, error) 
}
