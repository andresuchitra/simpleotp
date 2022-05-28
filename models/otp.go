package models

type OTPItem struct {
	ID       string `gorm:"primaryKey,uniqueIndex,not null" json:"id"`
	OTP      string `json:"otp" gorm:"index"`
	IsUsed   bool   `json:"-"`
	Phone    string `gorm:"not null" json:"phone" binding:"required" form:"phone"`
	ExpiryAt int64  `gorm:"not null" json:"expiry_at"`
}

type ValidateOTPParam struct {
	OTP   string `json:"otp" binding:"required"`
	Phone string `json:"phone" binding:"required"`
}
