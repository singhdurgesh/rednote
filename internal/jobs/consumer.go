package jobs

import (
	"github.com/singhdurgesh/rednote/internal/pkg/logger"
	queueService "github.com/singhdurgesh/rednote/internal/pkg/queue_service"
)

func ConsumerStart() {
	// Start the number of workers defined in configuration
	qs, err := queueService.NewRabbitMQService()

	qs.StartWorkers()

	if err != nil {
		logger.LogrusLogger.Fatalf("Error while starting the Queue Servier %v", err)
	}
}
