package models

type OTPItem struct {
	ID       string `gorm:"primaryKey,uniqueIndex,not null" json:"id"`
	OTP      string `json:"otp"`
	Phone    string `gorm:"not null" json:"phone" bindin:"required" form:"phone"`
	ExpiryAt int64  `gorm:"not null" json:"expiry_at"`
}
