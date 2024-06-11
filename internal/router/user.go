package router

import (
	"github.com/singhdurgesh/rednote/internal/app/controllers"
	"github.com/singhdurgesh/rednote/internal/middlewares"

	"github.com/gin-gonic/gin"
)

var userController = new(controllers.UserController)
var authController = new(controllers.AuthController)

func LoadUserRoutes(r *gin.Engine) *gin.RouterGroup {

	auth := r.Group("/users")
	{
		auth.POST("/signupByUsernamePassword", authController.SignupByUsernamePassword)
		auth.POST("/loginByUsernamePassword", authController.LoginByUsernamePassword)
		auth.POST("/auth/sendOtpPhone", authController.SendLoginOtpPhone)
		auth.POST("/auth/verifyOtpPhone", authController.VerifyLoginOtpPhone)
		auth.POST("/auth/resendOtpPhone", authController.ResendLoginOtpPhone)
	}

	user := r.Group("/users")
	user.Use(middlewares.Jwt())

	{
		user.GET("/profile", userController.GetUserProfile)
	}

	return auth
}
