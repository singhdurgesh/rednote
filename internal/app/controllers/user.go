package controllers

import (
	"net/http"

	"github.com/singhdurgesh/rednote/internal/app/services"

	"github.com/gin-gonic/gin"
)

var userService = new(services.UserService)

type UserController struct{}

type LoginByUsernamePasswordRequest struct {
	Usernmae string `json:"username" default:"admin"`
	Password string `json:"password" default:"123456"`
}

// @Router /users/loginByUsernamePassword [post]
// @Description Login By Username Password
// @Tags User
// @Param data body LoginByUsernamePasswordRequest true "username、password"
func (userController *UserController) LoginByUsernamePassword(ctx *gin.Context) {

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

	token := userService.LoginByUsernamePassword(username, password)
	if token == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Username or Password Error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"auth_token": token})
}

func (userController *UserController) SignupByUsernamePassword(ctx *gin.Context) {
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

	token, user := userService.SignupByUsernamePassword(data)

	if token == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "param error"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"token": token, "user": user})
}
