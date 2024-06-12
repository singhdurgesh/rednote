package services

import (
	"log"

	"github.com/singhdurgesh/rednote/cmd/app"
	"github.com/singhdurgesh/rednote/internal/app/models"
	"github.com/singhdurgesh/rednote/internal/app/services/otp_handler"
	"github.com/singhdurgesh/rednote/internal/constants"
	"github.com/singhdurgesh/rednote/internal/pkg/utils"
	"github.com/singhdurgesh/rednote/internal/tasks"
	"github.com/singhdurgesh/rednote/internal/tasks/notifications"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{}

var userService = new(UserService)

// @name LoginByUsernamePassword
// @description LoginByUsernamePassword
// @return string
func (authService *AuthService) LoginByUsernamePassword(username string, password string) (string, *models.User) {
	user := models.User{}

	res := app.Db.First(&user, "username = ?", username)

	if res.Error != nil || res.RowsAffected == 0 {
		return "", nil
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(password))
	if err != nil {
		return "", nil
	}

	token := authService.GenerateJwtToken(&user, "OTP")
	return token, &user
}

func (authService *AuthService) SignupByUsernamePassword(data map[string]interface{}) (string, *models.User) {
	userData := map[string]interface{}{"name": data["name"], "username": data["username"], "email": data["email"], "phone": data["phone"], "dob": data["dob"]}
	user := userService.CreateUser(userData)

	if user == nil {
		return "", user
	}

	err := userService.UpdatePassword(int(user.ID), data["password"].(string))

	if err != nil {
		log.Println(err.Error())
		return "", user
	}

	token := authService.GenerateJwtToken(user, "OTP")

	app.Db.Find(&user, "id = ?", user.ID)
	return token, user
}

func (authService *AuthService) SendLoginOtpPhone(phone string) (bool, string) {
	user := models.User{}
	user.Phone.Scan(phone)

	res := app.Db.Where(&user).FirstOrCreate(&user)

	if res.Error != nil {
		log.Println(res.Error, "Affected Rows: ", res.RowsAffected)
		return false, "Could not create the user"
	}

	// Generate and Send OTP Code
	task := notifications.NewLoginOtpCommunication(user)
	err := tasks.RunAsync(task)

	if err != nil {
		return false, err.Error()
	}

	return true, ""
}

func (authService *AuthService) ReSendLoginOtpPhone(phone string) (bool, string) {
	user := models.User{}
	user.Phone.Scan(phone)

	res := app.Db.Where(&user).FirstOrCreate(&user)

	if res.Error != nil {
		log.Println(res.Error, "Affected Rows: ", res.RowsAffected)
		return false, "Could not create the user"
	}

	// Generate and Send OTP Code
	task := notifications.NewResendLoginOtpCommunication(user)
	err := tasks.RunAsync(task)

	if err != nil {
		return false, err.Error()
	}

	return true, ""
}

func (authService *AuthService) VerifyLoginOtpPhone(phone string, otp string) (string, *models.User) {
	user := models.User{}
	res := app.Db.Find(&user, "phone = ?", phone)

	if res.Error != nil {
		log.Println(res.Error, "Affected Rows: ", res.RowsAffected)
		return "", nil
	}

	// Verify OTP Code
	result := otp_handler.ValidateOTP(phone, otp)

	if !result {
		return "", nil
	}

	token := authService.GenerateJwtToken(&user, constants.OTP_AUTH_MODE)

	return token, &user
}

func (authService *AuthService) GenerateJwtToken(user *models.User, authMode string) string {
	claims := utils.Claims{
		Username: user.Username.String,
		Uid:      user.ID,
		AuthMode: authMode,
	}

	return utils.GenerateToken(&claims)
}
