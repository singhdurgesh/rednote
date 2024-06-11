package controllers

import (
	"net/http"

	"github.com/singhdurgesh/rednote/internal/app/models"
	"github.com/singhdurgesh/rednote/internal/app/services"

	"github.com/gin-gonic/gin"
)

var userService = new(services.UserService)

type UserController struct{}

func (userController *UserController) GetUserProfile(ctx *gin.Context) {
	user, _ := ctx.Get("currentUser")

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

var updateUserParams = []string{"name", "dob", "email"}

func (userController *UserController) UpdateUser(ctx *gin.Context) {
	userData, _ := ctx.Get("currentUser")
	user := userData.(*models.User)

	data := make(map[string]interface{})

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	userParams := make(map[string]interface{})

	for _, attr := range updateUserParams {
		if val, ok := data[attr]; ok {
			userParams[attr] = val
		}
	}

	userData, err := userService.UpdateUser(int(user.ID), userParams)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	ctx.JSON(http.StatusOK, gin.H{"user": userData})
}
