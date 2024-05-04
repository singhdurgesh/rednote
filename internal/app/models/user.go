package models

import "time"

type User struct {
	BasicModel

	Name               string    `json:"name"`                        // Name
	Username           string    `json:"username" gorm:"uniqueIndex"` // Username
	Password           string    `json:"password"`                    // Password
	Phone              string    `json:"phone" gorm:"uniqueIndex"`    // Phone
	Email              string    `json:"email" gorm:"uniqueIndex"`    // Email
	OtpSecret          string    `json:"-"`
	ResetPasswordToken string    `json:"-"`
	LastLoginAt        time.Time `json:"last_login_at"` // Last login time
}
