// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/andresuchitra/simpleotp/models"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CreateOTP provides a mock function with given fields: ctx, newOtp
func (_m *Repository) CreateOTP(ctx *context.Context, newOtp *models.OTPItem) error {
	ret := _m.Called(ctx, newOtp)

	var r0 error
	if rf, ok := ret.Get(0).(func(*context.Context, *models.OTPItem) error); ok {
		r0 = rf(ctx, newOtp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRepository creates a new instance of Repository. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t testing.TB) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}