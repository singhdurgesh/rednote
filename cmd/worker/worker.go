package worker

import (
	"github.com/singhdurgesh/rednote/cmd/app"
	"github.com/singhdurgesh/rednote/configs"
	"github.com/singhdurgesh/rednote/internal/pkg/amqp"
	"github.com/singhdurgesh/rednote/internal/pkg/logger"
	"github.com/singhdurgesh/rednote/internal/pkg/postgres"
	"github.com/singhdurgesh/rednote/internal/pkg/redis"
	"github.com/singhdurgesh/rednote/internal/tasks/task_register"
)

func Init() {
	app.Config = configs.LoadConfig() // Setup Configuration

	// Connect Logger
	app.Logger = logger.Init()

	// connect Database
	app.Db = postgres.Connect(&app.Config.Postgres)

	// Connect Cache
	app.Cache = redis.Connect(&app.Config.Redis)

	// Start Work Task Server
	app.Broker = amqp.Connect(&app.Config.AMQPConfig)
	task_register.RegisterTasks()

	amqp.StartWorker(app.Broker, app.Config.AMQPConfig.WorkerPoolSize)
}
