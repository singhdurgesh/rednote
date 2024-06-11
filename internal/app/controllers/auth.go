package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/singhdurgesh/rednote/internal/app/services"
)

type AuthController struct{}

var authService = new(services.AuthService)

func (authController *AuthController) LoginByUsernamePassword(ctx *gin.Context) {

	data := make(map[string]interface{})

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}
	username := data["username"].(string)
	password := data["password"].(string)

	if username == "" || password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "error param"})
		return
	}

	token, user := authService.LoginByUsernamePassword(username, password)
	if token == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Username or Password Invalid"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token, "user": user})
}

func (authController *AuthController) SignupByUsernamePassword(ctx *gin.Context) {
	data := make(map[string]interface{})

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	password := data["password"]
	password_confirmation := data["password_confirmation"]
	username := data["username"]

	if password == "" || password == nil || password_confirmation == "" || password_confirmation == nil || username == "" || username == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "username/password should be present"})
		return
	}

	if password != password_confirmation {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "passwords didn't match"})
		return
	}

	token, user := authService.SignupByUsernamePassword(data)

	if token == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "param error"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"token": token, "user": user})
}

func (authController *AuthController) SendLoginOtpPhone(ctx *gin.Context) {
	data := make(map[string]interface{})

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}
	phone := data["phone"]

	if phone == nil || phone == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "error param"})
		return
	}

	result, err := authService.SendLoginOtpPhone(data["phone"].(string))

	if !result {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "successful"})
}

func (authController *AuthController) VerifyLoginOtpPhone(ctx *gin.Context) {
	data := make(map[string]interface{})

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}
	phone := data["phone"]
	otp := data["otp"]

	if phone == nil || phone == "" || otp == nil || otp == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "error param"})
		return
	}

	token, user := authService.VerifyLoginOtpPhone(data["phone"].(string), data["otp"].(string))

	if token == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Invalid Mobile or OTP"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token, "user": user})
}

func (authController *AuthController) ResendLoginOtpPhone(ctx *gin.Context) {
	data := make(map[string]interface{})

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}
	phone := data["phone"]

	if phone == nil || phone == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "error param"})
		return
	}

	result, err := authService.ReSendLoginOtpPhone(data["phone"].(string))

	if !result {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "successful"})
}
