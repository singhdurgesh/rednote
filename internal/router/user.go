package router

import (
	"github.com/singhdurgesh/rednote/internal/controllers"

	"github.com/gin-gonic/gin"
)

var userController = new(controllers.UserController)

func LoadUserRoutes(r *gin.Engine) *gin.RouterGroup {

	user := r.Group("/users")
	{
		user.POST("/loginByUsernamePassword", userController.LoginByUsernamePassword)
	}
	return user
}
