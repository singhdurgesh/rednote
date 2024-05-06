package queueService

import (
	"time"
)

type QueueService interface {
	CreateQueue(string, bool) error
	DeleteQueue(string, bool, bool, bool) error
	ClearQueue(string, bool) error
	PushJob(string, []byte, string, time.Time) error
}
