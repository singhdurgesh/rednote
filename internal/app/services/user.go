package services

import (
	"log"

	"github.com/singhdurgesh/rednote/cmd/app"
	"github.com/singhdurgesh/rednote/internal/app/models"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

func (userService *UserService) CreateUser(data map[string]interface{}) *models.User {
	res := app.Db.Model(&models.User{}).Create(data)

	if res.Error != nil || res.RowsAffected == 0 {
		log.Println(res.Error, "Affected Rows: ", res.RowsAffected)
		return nil
	}

	user := models.User{}
	app.Db.Find(&user, "id = ?", data["id"])

	return &user
}

func (userService *UserService) UpdatePassword(userId int, password string) error {
	user := models.User{}
	user.ID = uint(userId)

	password_hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return err
	}
	data := map[string]interface{}{"password": password_hash}

	res := app.Db.Model(&user).Updates(data)

	if res.Error != nil || res.RowsAffected == 0 {
		return res.Error
	}
	return nil
}

func (userService *UserService) GetActiveUserById(userId int) *models.User {
	user := models.User{}
	res := app.Db.Find(&user, "id = ?", userId)

	if res.Error != nil {
		log.Println(res.Error)
		return nil
	}

	return &user
}
