package router

import (
	"github.com/singhdurgesh/rednote/internal/app/controllers"

	"github.com/gin-gonic/gin"
)

var userController = new(controllers.UserController)

func LoadUserRoutes(r *gin.Engine) *gin.RouterGroup {

	user := r.Group("/users")
	{
		user.POST("/signupByUsernamePassword", userController.SignupByUsernamePassword)
		user.POST("/loginByUsernamePassword", userController.LoginByUsernamePassword)
		user.POST("/auth/sendOtpPhone", userController.SendLoginOtpPhone)
		user.POST("/auth/verifyOtpPhone", userController.VerifyLoginOtpPhone)
		user.POST("/auth/resendOtpPhone", userController.ResendLoginOtpPhone)
	}
	return user
}
