package controllers

import (
	"net/http"

	"github.com/singhdurgesh/rednote/internal/app/services"

	"github.com/gin-gonic/gin"
)

var userService = new(services.UserService)

type UserController struct{}

func (userController *UserController) GetUserProfile(ctx *gin.Context) {
	user, _ := ctx.Get("currentUser")

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}
