package postgres

import (
	"fmt"

	"github.com/singhdurgesh/rednote/configs"
	"github.com/singhdurgesh/rednote/internal/models"
	"github.com/singhdurgesh/rednote/internal/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(config *configs.Postgres) *gorm.DB {

	logger := logger.LogrusLogger

	address := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Kolkata",
		config.Host,
		config.Username,
		config.Password,
		config.Database,
		config.Port,
	)

	fmt.Println(address)

	// refer https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL for details
	db, err := gorm.Open(postgres.Open(address), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic(`😫: Connected failed, check your Postgres with ` + address)
	}

	// Migrate the schema
	migrateErr := db.AutoMigrate(&models.Example{}, &models.User{})
	if migrateErr != nil {
		panic(`😫: Auto migrate failed, check your Postgres with ` + address)
	}

	// export DB
	DB = db

	logger.Printf("🍟: Successfully connected to Postgres at "+address, DB)

	return db

}
