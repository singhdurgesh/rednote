package router

import (
	"github.com/singhdurgesh/rednote/internal/app/controllers/adminControllers"

	"github.com/gin-gonic/gin"
)

var authController = new(adminControllers.AuthController)

func LoadAdminRoutes(r *gin.Engine) *gin.RouterGroup {
	admin := r.Group("/admin")

	{
		admin.POST("/googleLogin", authController.GoogleLogin)
	}

	return admin
}
