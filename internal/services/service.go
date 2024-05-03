package services

import (
	"github.com/singhdurgesh/rednote/internal/pkg/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	db = postgres.DB
}
