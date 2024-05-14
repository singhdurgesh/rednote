package rabbitmq

import (
	"fmt"

	"github.com/singhdurgesh/rednote/configs"
	"github.com/streadway/amqp"
)

func Connect() (*amqp.Connection, error) {
	rc := configs.EnvConfig.Rabbitmq

	connAddress := fmt.Sprintf(
		"amqp://%s:%s@%s:%s",
		rc.User,
		rc.Password,
		rc.Host,
		rc.Port,
	)

	return amqp.Dial(connAddress)
}
