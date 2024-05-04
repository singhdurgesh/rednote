package models

import (
	"time"

	"gorm.io/gorm"
)

type BasicModel struct {
	*gorm.Model
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
