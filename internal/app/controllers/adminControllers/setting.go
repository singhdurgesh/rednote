package adminControllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/singhdurgesh/rednote/cmd/app"
	"github.com/singhdurgesh/rednote/internal/app/models"
	settings "github.com/singhdurgesh/rednote/internal/app/services/settingsService"
)

type SettingController struct{}

var settingService = new(settings.SettingService)

func (settingController *SettingController) GetAll(ctx *gin.Context) {
	settings, err := settingService.GetAll()

	if err != nil {
		app.Logger.Error(err)
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"settings": settings})
}

func (settingController *SettingController) CreateSetting(ctx *gin.Context) {
	setting := &models.Setting{}

	if err := ctx.ShouldBindJSON(&setting); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	err := settingService.CreateSetting(setting)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"setting": setting})
}

func (settingController *SettingController) GetSetting(ctx *gin.Context) {
	settingKey := ctx.Param("settingKey")
	app.Logger.Println(settingKey)

	setting, err := settingService.GetSettingByKey(settingKey)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"setting": setting})
}

var updateSettingParams = []string{"value", "title", "description"}

func (settingController *SettingController) UpdateSetting(ctx *gin.Context) {
	settingKey := ctx.Param("settingKey")
	setting, err := settingService.GetSettingByKey(settingKey)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	data := make(map[string]interface{})

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	for _, attr := range updateSettingParams {
		if data[attr] != "" {
			switch attr {
			case "value":
				setting.Value = data[attr].(string)
			case "title":
				setting.Title = data[attr].(string)
			case "description":
				setting.Description = data[attr].(string)
			}
		}
	}

	resp := app.Db.Model(&setting).Updates(setting)

	if resp.Error != nil || resp.RowsAffected == 0 {
		ctx.JSON(http.StatusBadRequest, "Couldn't update the setting")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"setting": setting})
}

func (settingController *SettingController) DeleteSetting(ctx *gin.Context) {
	settingKey := ctx.Param("settingKey")
	setting, err := settingService.GetSettingByKey(settingKey)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	resp := app.Db.Model(&setting).Delete(&setting)

	if resp.Error != nil || resp.RowsAffected == 0 {
		ctx.JSON(http.StatusBadRequest, "Couldn't delete the setting")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "Sucessfully Deleted the Setting"})
}
