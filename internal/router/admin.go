package router

import (
	"github.com/singhdurgesh/rednote/internal/app/controllers/adminControllers"
	"github.com/singhdurgesh/rednote/internal/middlewares"

	"github.com/gin-gonic/gin"
)

var adminAuthController, settingController = new(adminControllers.AuthController), new(adminControllers.SettingController)

func LoadAdminRoutes(r *gin.Engine) *gin.RouterGroup {
	admin := r.Group("/admin")

	{
		admin.POST("/googleLogin", adminAuthController.GoogleLogin)
	}

	setting := admin.Group("/settings")
	setting.Use(middlewares.Jwt())

	{
		setting.GET("/", settingController.GetAll)
		setting.POST("/", settingController.CreateSetting)
		setting.GET("/:settingKey", settingController.GetSetting)
		setting.PUT("/:settingKey", settingController.UpdateSetting)
		setting.DELETE("/:settingKey", settingController.DeleteSetting)
	}

	return admin
}
