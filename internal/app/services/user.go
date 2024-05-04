package services

import (
	"log"

	"github.com/singhdurgesh/rednote/internal/app/models"
	"github.com/singhdurgesh/rednote/internal/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

// @name LoginByUsernamePassword
// @description LoginByUsernamePassword
// @return string
func (userService *UserService) LoginByUsernamePassword(username string, password string) string {

	user := models.User{}

	res := db.First(&user, "username = ?", username)

	if res.Error != nil || res.RowsAffected == 0 {
		return ""
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(password))
	if err != nil {
		return ""
	}

	token := userService.GenerateJwtToken(&user)
	return token
}

func (userService *UserService) SignupByUsernamePassword(data map[string]interface{}) (string, *models.User) {
	userData := map[string]interface{}{"name": data["name"], "username": data["username"], "email": data["email"], "phone": data["phone"]}
	user := userService.CreateUser(userData)

	if user == nil {
		return "", user
	}

	err := userService.UpdatePassword(int(user.ID), data["password"].(string))

	if err != nil {
		log.Println(err.Error())
		return "", user
	}

	token := userService.GenerateJwtToken(user)

	db.Find(&user, "id = ?", user.ID)
	return token, user
}

func (userService *UserService) CreateUser(data map[string]interface{}) *models.User {
	res := db.Model(&models.User{}).Create(data)

	// res := db.Create(&user)

	if res.Error != nil || res.RowsAffected == 0 {
		log.Println(res.Error, "Affected Rows: ", res.RowsAffected)
		return nil
	}

	user := models.User{}
	db.Find(&user, "id = ?", data["id"])

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

	res := db.Model(&user).Updates(data)

	if res.Error != nil || res.RowsAffected == 0 {
		return res.Error
	}
	return nil
}

func (userService *UserService) GenerateJwtToken(user *models.User) string {
	claims := utils.Claims{
		Username: user.Username.String,
		Uid:      user.ID,
	}

	return utils.GenerateToken(&claims)
}
