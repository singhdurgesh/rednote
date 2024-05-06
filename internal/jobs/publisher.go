package jobs

import (
	"time"

	"github.com/singhdurgesh/rednote/internal/pkg/logger"
	queueService "github.com/singhdurgesh/rednote/internal/pkg/queue_service"
)

var Publisher *JobsPublisher

func PublisherStart() {
	qs, err := queueService.NewRabbitMQService()

	if err != nil {
		logger.LogrusLogger.Fatalf("Error while starting the Queue Servier %v", err)
	}

	Publisher = NewJobsPublisher(qs)
}

type JobsPublisher struct {
	queueService queueService.QueueService
}

func NewJobsPublisher(queueService queueService.QueueService) *JobsPublisher {
	return &JobsPublisher{queueService: queueService}
}

func (j *JobsPublisher) Publish(queue string, message string, contentType string) {
	j.queueService.PushJob(queue, []byte(message), contentType, time.Now())
}

func (j *JobsPublisher) PublishwithDelay(queue string, message string, contentType string, time time.Time) {
	j.queueService.PushJob(queue, []byte(message), contentType, time)
}
