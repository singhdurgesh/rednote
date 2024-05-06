package queueService

import (
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/singhdurgesh/rednote/internal/pkg/rabbitmq"
	"github.com/streadway/amqp"
)

type RabbitMQService struct {
	Amqpchan *amqp.Channel
}

func NewRabbitMQService() (*RabbitMQService, error) {
	rs := RabbitMQService{}
	err := rs.Init()

	return &rs, err
}

func (r *RabbitMQService) Init() error {
	aC, err := r.GetChannel()
	r.Amqpchan = aC

	return err
}

func (r *RabbitMQService) CreateQueue(queue string, durable bool) error {
	_, err := r.Amqpchan.QueueDeclare(queue, durable, false, false, false, nil)

	return err
}

func (r *RabbitMQService) ClearQueue(queue string, noWait bool) error {
	_, err := r.Amqpchan.QueuePurge(queue, noWait)

	return err
}

func (r *RabbitMQService) DeleteQueue(queue string, unUsed, ifEmpty, noWait bool) error {
	_, err := r.Amqpchan.QueueDelete(queue, unUsed, ifEmpty, noWait)

	return err
}

func (r *RabbitMQService) PushJob(queue string, message []byte, contentType string, executionTime time.Time) error {
	if err := r.Amqpchan.Publish(
		"",    // exchange_key
		queue, // routing key
		false, // exclusive
		false, // mandatory
		amqp.Publishing{
			ContentType:  contentType,
			DeliveryMode: amqp.Transient,
			MessageId:    uuid.New().String(),
			Timestamp:    executionTime,
			Body:         message,
		},
	); err != nil {
		return errors.Wrap(err, "ch.Publish")
	}

	return nil
}

func (r *RabbitMQService) GetChannel() (*amqp.Channel, error) {
	conn, err := rabbitmq.Connect()

	if err != nil {
		return nil, err
	}

	return conn.Channel()
}
