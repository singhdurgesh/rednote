package services

import (
	"github.com/singhdurgesh/rednote/internal/models"
	"github.com/singhdurgesh/rednote/internal/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

// @name LoginByUsernamePassword
// @description LoginByUsernamePassword
// @return string
func (userService *UserService) LoginByUsernamePassword(username string, password string) string {

	user := models.User{
		Username: username,
		Password: password,
		BasicModel: models.BasicModel{
			ID: 5,
		},
	}

	res := db.First(&user, "username = ?", username)

	if res.Error != nil || res.RowsAffected == 0 {
		return ""
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return ""
	}
	claims := utils.Claims{
		Username: user.Username,
		Uid:      user.ID,
	}
	token := utils.GenerateToken(&claims)

	return token
}
