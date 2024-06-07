package services

import (
	"errors"

	"github.com/singhdurgesh/rednote/cmd/app"
	"github.com/singhdurgesh/rednote/internal/app/models"
)

type ExampleService struct{}

func (exampleService *ExampleService) CreateExample(data map[string]interface{}) (*models.Example, error) {

	name := ""

	if data["name"] != nil {
		name = data["name"].(string)
	}

	if name == "" {
		return nil, errors.New("name should be present")
	}

	example := models.Example{
		Name: name,
	}

	res := app.Db.Create(&example)
	if res.Error != nil || res.RowsAffected == 0 {
		return nil, res.Error
	}

	return &example, nil
}

func (exampleService *ExampleService) GetExample(exampleId int) *models.Example {

	example := models.Example{}
	res := app.Db.First(&example, exampleId).Select("id, name, status")

	if res.Error != nil || res.RowsAffected == 0 {
		return nil
	}
	return &example
}

func (exampleService *ExampleService) UpdateExample(data map[string]interface{}) bool {

	exampleId := ""

	if data["exampleId"] != nil {
		exampleId = data["exampleId"].(string)
	}

	if exampleId == "" {
		return false
	}

	example := models.Example{}
	example.ID = uint(data["exampleId"].(int))

	res := app.Db.Model(&example).Updates(data)

	if res.Error != nil || res.RowsAffected == 0 {
		return false
	}
	return true
}

func (exampleService *ExampleService) DeleteExample(exampleId int) bool {

	example := models.Example{}
	res := app.Db.Delete(&example, exampleId)

	if res.Error != nil || res.RowsAffected == 0 {
		return false
	}
	return true
}
