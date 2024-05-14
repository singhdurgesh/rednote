package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/singhdurgesh/rednote/internal/tasks"
	"github.com/singhdurgesh/rednote/internal/tasks/notifications"
)

type JobQueueController struct{}

func (JobQueueController *JobQueueController) PushJob(ctx *gin.Context) {
	data := make(map[string]interface{})
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	if data["content"] == nil || data["content"] == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "content should be present"})
	}

	task := notifications.NewNotificationTask("9721323", 1234, "Random Data")

	err := tasks.RunAsync(task)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "Jobs Pushed"})
}

// WIP: Not Ready for Use
func (JobQueueController *JobQueueController) PushJobDelay(ctx *gin.Context) {
	data := make(map[string]interface{})
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	if data["content"] == nil || data["content"] == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "content should be present"})
	}

	// content := data["content"].(string)

	// queue := "jobs"
	// content_type := "application/json"

	// if data["content_type"] != nil {
	// 	content_type = data["content_type"].(string)
	// }

	// if data["queue"] != nil {
	// 	queue = data["queue"].(string)
	// }

	// delayInMinutes := 0.0

	// if data["delayInMinutes"] != nil {
	// 	delayInMinutes = data["delayInMinutes"].(float64)
	// }

	ctx.JSON(http.StatusCreated, gin.H{"status": "Jobs Pushed"})
}
