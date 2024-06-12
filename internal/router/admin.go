package router

import (
	"github.com/singhdurgesh/rednote/internal/app/controllers/adminControllers"

	"github.com/gin-gonic/gin"
)

var adminAuthController = new(adminControllers.AuthController)

func LoadAdminRoutes(r *gin.Engine) *gin.RouterGroup {
	admin := r.Group("/admin")

	{
		admin.POST("/googleLogin", adminAuthController.GoogleLogin)
	}

	return admin
}
