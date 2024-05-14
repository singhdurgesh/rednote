package router

import (
	"github.com/singhdurgesh/rednote/internal/app/controllers"
	"github.com/singhdurgesh/rednote/internal/middlewares"

	"github.com/gin-gonic/gin"
)

var jobQueueController = new(controllers.JobQueueController)

func LoadJobQueueRoutes(r *gin.Engine) *gin.RouterGroup {

	job_queue := r.Group("/jobs/")
	job_queue.Use(middlewares.Jwt())

	{
		job_queue.POST("/pushJob", jobQueueController.PushJob)
		job_queue.POST("/pushJobDelay", jobQueueController.PushJobDelay)
	}
	return job_queue
}
