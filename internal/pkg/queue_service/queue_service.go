package queueService

import (
	"time"

	"github.com/singhdurgesh/rednote/internal/pkg/logger"
)

type QueueService interface {
	CreateQueue(string, bool) error
	DeleteQueue(string, bool, bool, bool) error
	ClearQueue(string, bool) error
	PushJob(string, []byte, string, time.Time) error
	StartWorkers() error
}

func ProcessJob(q QueueService, task []byte, task_type string) error {
	logger.LogrusLogger.Println(string(task), task_type, time.Now())
	return nil
}
