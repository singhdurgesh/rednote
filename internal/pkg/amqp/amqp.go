package amqp

import (
	"fmt"
	"log"

	"github.com/RichardKnop/machinery/v2"
	amqpbackend "github.com/RichardKnop/machinery/v2/backends/amqp"
	amqpbroker "github.com/RichardKnop/machinery/v2/brokers/amqp"
	"github.com/RichardKnop/machinery/v2/config"
	"github.com/RichardKnop/machinery/v2/locks/eager"
	"github.com/RichardKnop/machinery/v2/tasks"
	"github.com/singhdurgesh/rednote/configs"
)

func Connect(AMQPConfig *configs.AMQPConfig) *machinery.Server {
	c := &config.Config{
		DefaultQueue:    AMQPConfig.Queue,
		ResultsExpireIn: 3600,
		ResultBackend: fmt.Sprintf(
			"%s://%s:%s@%s:%s/",
			AMQPConfig.Protocol,
			AMQPConfig.User,
			AMQPConfig.Password,
			AMQPConfig.Host,
			AMQPConfig.Port,
		),
		Broker: fmt.Sprintf(
			"%s://%s:%s@%s:%s/",
			AMQPConfig.Protocol,
			AMQPConfig.User,
			AMQPConfig.Password,
			AMQPConfig.Host,
			AMQPConfig.Port,
		),
		AMQP: &config.AMQPConfig{
			Exchange:      AMQPConfig.Exchange,
			ExchangeType:  AMQPConfig.ExchangeType,
			BindingKey:    AMQPConfig.BindingKey,
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
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		panic(err)
	}

	task := tasks.Signature{
		Name: "task",
		Args: []tasks.Arg{
			{
				Type:  "string",
				Value: "payload",
			},
		},
	}
	_, err = server.SendTask(&task)

	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		panic(err)
	}

	return server
}

func StartWorker(Server *machinery.Server, WorkerCount int) error {
	worker := Server.NewWorker("", WorkerCount)
	if err := worker.Launch(); err != nil {
		return err
	}

	return nil
}
