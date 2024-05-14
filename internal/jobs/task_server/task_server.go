package task_server

import (
	"fmt"

	"github.com/RichardKnop/machinery/v2"
	amqpbackend "github.com/RichardKnop/machinery/v2/backends/amqp"
	amqpbroker "github.com/RichardKnop/machinery/v2/brokers/amqp"
	"github.com/RichardKnop/machinery/v2/config"
	"github.com/RichardKnop/machinery/v2/locks/eager"
	"github.com/singhdurgesh/rednote/configs"
	"github.com/singhdurgesh/rednote/internal/jobs/tasks/notification"
	"github.com/singhdurgesh/rednote/internal/pkg/logger"
)

var Runner *TaskRunner

type TaskRunner struct {
	Server *machinery.Server
}

func StartServer() {
	NewTaskRunner()
}

func NewTaskRunner() *TaskRunner {
	c := &config.Config{
		DefaultQueue:    "machinery_tasks",
		ResultsExpireIn: 3600,
		ResultBackend: fmt.Sprintf(
			"amqp://%s:%s@%s:%s/",
			configs.EnvConfig.Rabbitmq.User,
			configs.EnvConfig.Rabbitmq.Password,
			configs.EnvConfig.Rabbitmq.Host,
			configs.EnvConfig.Rabbitmq.Port,
		),
		Broker: fmt.Sprintf(
			"amqp://%s:%s@%s:%s/",
			configs.EnvConfig.Rabbitmq.User,
			configs.EnvConfig.Rabbitmq.Password,
			configs.EnvConfig.Rabbitmq.Host,
			configs.EnvConfig.Rabbitmq.Port,
		),
		AMQP: &config.AMQPConfig{
			Exchange:      "machinery_exchange",
			ExchangeType:  "direct",
			BindingKey:    "machinery_task",
			PrefetchCount: 3,
		},
	}
	broker := amqpbroker.New(c)
	backend := amqpbackend.New(c)
	lock := eager.New()

	server := machinery.NewServer(c, broker, backend, lock)
	err := server.RegisterTasks(map[string]interface{}{
		"task": func() error {
			fmt.Println("Task Done")
			return nil
		},
	})
	if err != nil {
		panic(err)
	}

	t := &TaskRunner{
		Server: server,
	}

	Runner = t

	return t
}

func StartWorker() error {
	worker := Runner.Server.NewWorker("", configs.EnvConfig.Rabbitmq.WorkerPoolSize)
	if err := worker.Launch(); err != nil {
		return err
	}

	return nil
}

func RegisterTasks() {
	// Register New Tasks here
	err := Runner.Server.RegisterTasks(map[string]interface{}{
		"send_otp": notification.SendMail,
		// "notification": tasks.ProcessNotification,
	})

	if err != nil {
		logger.LogrusLogger.Fatalln(err)
	}
}
