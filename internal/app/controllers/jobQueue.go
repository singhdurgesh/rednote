package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/singhdurgesh/rednote/internal/jobs"
	"github.com/singhdurgesh/rednote/internal/jobs/tasks"
	queueService "github.com/singhdurgesh/rednote/internal/pkg/queue_service"
)

type JobQueueController struct{}

func (JobQueueController *JobQueueController) CreateQueue(ctx *gin.Context) {
	data := make(map[string]interface{})
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	// err := jobs.Publisher.CreateQueue(data["queue"].(string), data["durable"].(bool))
	qs, err := queueService.NewRabbitMQService()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	if data["queue"] == nil || data["queue"] == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "queue name should be present"})
	}

	queue_name := data["queue"]

	durable := false

	if data["durable"] != nil {
		durable, _ = strconv.ParseBool(data["durable"].(string))
	}

	if err = qs.CreateQueue(queue_name.(string), durable); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "Queue Created"})
}

func (JobQueueController *JobQueueController) DeleteQueue(ctx *gin.Context) {
	data := make(map[string]interface{})
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	// err := jobs.Publisher.CreateQueue(data["queue"].(string), data["durable"].(bool))
	qs, err := queueService.NewRabbitMQService()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	if data["queue"] == nil || data["queue"] == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "queue name should be present"})
	}

	queue_name := data["queue"]

	nowait := false
	unused := true
	ifempty := true

	if data["nowait"] != nil {
		nowait, _ = strconv.ParseBool(data["nowait"].(string))
	}

	if data["unused"] != nil {
		unused, _ = strconv.ParseBool(data["unused"].(string))
	}

	if data["ifempty"] != nil {
		unused, _ = strconv.ParseBool(data["ifempty"].(string))
	}

	if err = qs.DeleteQueue(queue_name.(string), unused, ifempty, nowait); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "Queue Deleted"})
}

func (JobQueueController *JobQueueController) PushJob(ctx *gin.Context) {
	data := make(map[string]interface{})
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	if data["content"] == nil || data["content"] == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "content should be present"})
	}

	err := tasks.NewNotificationTask("9721323", 1234, "Random Data").RunAsync()

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

	content := data["content"].(string)

	queue := "jobs"
	content_type := "application/json"

	if data["content_type"] != nil {
		content_type = data["content_type"].(string)
	}

	if data["queue"] != nil {
		queue = data["queue"].(string)
	}

	delayInMinutes := 0.0

	if data["delayInMinutes"] != nil {
		delayInMinutes = data["delayInMinutes"].(float64)
	}

	if err := jobs.Publisher.PublishwithDelay(queue, content, content_type, time.Now().Add(time.Minute*time.Duration(delayInMinutes))); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "Jobs Pushed"})
}
