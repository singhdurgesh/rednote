package queueService

import (
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/singhdurgesh/rednote/configs"
	"github.com/singhdurgesh/rednote/internal/pkg/logger"
	"github.com/singhdurgesh/rednote/internal/pkg/rabbitmq"
	"github.com/streadway/amqp"
)

const (
	queueAutoDelete = false
	queueExclusive  = false
	queueNoWait     = false

	publishMandatory = false
	publishImmediate = false

	prefetchCount  = 1
	prefetchSize   = 0
	prefetchGlobal = false

	consumeAutoAck   = false
	consumeExclusive = false
	consumeNoLocal   = false
	consumeNoWait    = false
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
	_, err := r.Amqpchan.QueueDeclare(queue, durable, queueAutoDelete, queueExclusive, queueNoWait, nil)

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
		configs.EnvConfig.Rabbitmq.Exchange, // exchange_key
		configs.EnvConfig.Rabbitmq.Queue,    // routing key
		publishMandatory,
		publishImmediate, //
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

func (r *RabbitMQService) StartWorkers() error {
	err := r.Amqpchan.Qos(
		prefetchCount,  // prefetch count
		prefetchSize,   // prefetch size
		prefetchGlobal, // global
	)

	if err != nil {
		return errors.Wrap(err, "Error  ch.Qos")
	}

	tasks, err := r.Amqpchan.Consume(
		configs.EnvConfig.Rabbitmq.Queue,
		configs.EnvConfig.Rabbitmq.ConsumerTag, // consumer tag
		consumeAutoAck,                         // consumerAutoAck
		consumeExclusive,                       // consumeExclusive
		consumeNoLocal,                         // ConsumeNoLocal
		consumeNoWait,                          // consumeNowait
		nil,
	)

	if err != nil {
		return errors.Wrap(err, "Consume")
	}

	for i := 0; i < configs.EnvConfig.Rabbitmq.WorkerPoolSize; i++ {
		go r.worker(tasks)
	}

	chanErr := <-r.Amqpchan.NotifyClose(make(chan *amqp.Error))
	logger.LogrusLogger.Errorf("ch.NotifyClose: %v", chanErr)
	return chanErr
}

func (r *RabbitMQService) worker(messages <-chan amqp.Delivery) {
	for task := range messages {
		err := ProcessJob(r, task.Body, task.ContentType)
		if err != nil {
			if err := task.Reject(false); err != nil {
				logger.LogrusLogger.Errorf("Err delivery.Reject: %v", err)
			}
			logger.LogrusLogger.Errorf("Failed to process delivery: %v", err)
		} else {
			err = task.Ack(false)
			if err != nil {
				logger.LogrusLogger.Errorf("Failed to acknowledge delivery: %v", err)
			}
		}
	}

	logger.LogrusLogger.Println("Tasks Channel Closed")
}
