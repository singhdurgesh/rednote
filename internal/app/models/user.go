package models

import (
	"time"

	"github.com/singhdurgesh/rednote/pkg/datatypes"
)

type User struct {
	BasicModel

	Name          datatypes.NullString `json:"name"`                        // Name
	Username      datatypes.NullString `json:"username" gorm:"uniqueIndex"` // Username
	Password      datatypes.NullString `json:"-"`                           // Password
	Phone         datatypes.NullString `json:"phone" gorm:"uniqueIndex"`    // Phone
	Email         datatypes.NullString `json:"email" gorm:"uniqueIndex"`    // Email
	EmailVerified bool                 `json:"email_verified"`
	LastLoginAt   time.Time            `json:"-"` // Last login time
}
