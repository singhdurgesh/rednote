package postgres

import (
	"fmt"
	"time"

	"github.com/singhdurgesh/rednote/cmd/app"
	"github.com/singhdurgesh/rednote/configs"
	"github.com/singhdurgesh/rednote/internal/app/models"
	"github.com/spf13/viper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gorm_logger "gorm.io/gorm/logger"
)

func Connect(config *configs.Postgres) *gorm.DB {
	address := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s sslrootcert=%s sslcert=%s sslkey=%s TimeZone=Asia/Kolkata",
		config.Host,
		config.Username,
		config.Password,
		config.Database,
		config.Port,
		config.Sslmode,
		config.Sslrootcert,
		config.Sslcert,
		config.Sslkey,
	)

	logLevel := gorm_logger.Info

	if viper.Get("env") == "staging" || viper.Get("env") == "production" {
		logLevel = gorm_logger.Warn
	}

	db_logger := gorm_logger.New(app.Logger,
		gorm_logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logLevel,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		},
	)

	// refer https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL for details
	db, err := gorm.Open(postgres.Open(address), &gorm.Config{
		Logger: db_logger,
	})

	if err != nil {
		fmt.Println(err)
		panic(`😫: Connected failed, check your Postgres with ` + address)
	}

	// Migrate the schema
	migrateErr := db.AutoMigrate(&models.Example{}, &models.User{})
	if migrateErr != nil {
		panic(`😫: Auto migrate failed, check your Postgres with ` + address)
	}

	app.Logger.Printf("🍟: Successfully connected to Postgres: %v", db)

	return db

}
