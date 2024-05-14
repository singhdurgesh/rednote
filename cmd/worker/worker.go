package worker

import (
	"github.com/singhdurgesh/rednote/configs"
	"github.com/singhdurgesh/rednote/internal/app/services"
	"github.com/singhdurgesh/rednote/internal/jobs/task_register"
	"github.com/singhdurgesh/rednote/internal/jobs/task_server"
	"github.com/singhdurgesh/rednote/internal/pkg/logger"
	"github.com/singhdurgesh/rednote/internal/pkg/postgres"
)

func Init() {
	configs.LoadConfig() // Setup Configuration

	logger.Init()

	// connect Database
	postgres.Connect(&configs.EnvConfig.Postgres)

	services.Init()

	task_server.StartServer()
	task_register.RegisterTasks()
	err := task_server.StartWorker()

	if err != nil {
		logger.LogrusLogger.Panic(err)
	}
}
