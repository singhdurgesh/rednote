package settingsService

import (
	"errors"
	"log"

	"github.com/singhdurgesh/rednote/cmd/app"
	"github.com/singhdurgesh/rednote/internal/app/models"
)

type SettingService struct{}

func GetValue(key string) string {
	setting := models.Setting{}

	resp := app.Db.Model(&models.Setting{}).First(&setting, "key = ?", key)

	if resp.Error != nil || resp.RowsAffected == 0 {
		return ""
	}

	return setting.Value
}

func (settingService *SettingService) GetAll() ([]models.Setting, error) {
	var settings []models.Setting

	res := app.Db.Find(&settings)

	if res.Error != nil {
		return settings, res.Error
	}

	return settings, nil
}

func (settingService *SettingService) CreateSetting(data *models.Setting) error {
	res := app.Db.Model(&models.Setting{}).Create(data)

	if res.Error != nil || res.RowsAffected == 0 {
		log.Println(res.Error, "Affected Rows: ", res.RowsAffected)
		return res.Error
	}

	return nil
}

func (settingService *SettingService) GetSettingByKey(settingKey string) (*models.Setting, error) {
	setting := models.Setting{}

	resp := app.Db.Model(&models.Setting{}).First(&setting, "key = ?", settingKey)

	if resp.Error != nil || resp.RowsAffected == 0 {
		return nil, errors.New("setting does not exists")
	}

	return &setting, nil
}
