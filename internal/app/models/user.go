package models

import (
	"time"

	"github.com/singhdurgesh/rednote/pkg/datatypes"
)

type User struct {
	BasicModel

	Name               datatypes.NullString // Name
	Username           datatypes.NullString `json:"username" gorm:"uniqueIndex"` // Username
	Password           datatypes.NullString `json:"-"`                           // Password
	Phone              datatypes.NullString `json:"phone" gorm:"uniqueIndex"`    // Phone
	Email              datatypes.NullString `json:"email" gorm:"uniqueIndex"`    // Email
	OtpSecret          string               `json:"-"`
	ResetPasswordToken string               `json:"-"`
	LastLoginAt        time.Time            `json:"-"` // Last login time
}
