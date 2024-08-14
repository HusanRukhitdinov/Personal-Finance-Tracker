package rabbitMq

import (
	"budgeting_service/storage"
	"log"
	"time"

	"github.com/streadway/amqp"
)

type ConsumerRabbitMq interface {
	ConsumeMessages(handler func(message []byte)) error
	Close()
}

type RabbitMqConsumerImpl struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	queueName string
	storage   storage.IStorage
}

func NewRabbitMqConsumerImpl(url string, queue string, storage storage.IStorage) (*RabbitMqConsumerImpl, error) {
	var err error
	var conn *amqp.Connection
	for i := 0; i < 10; i++ {
		conn, err = amqp.Dial(url)
		if err != nil {
			log.Println("Failed to connect to RabbitMQ")
			time.Sleep(1 * time.Second)
			continue
		}
	}

	// conn, err := amqp.Dial(url)
	// if err != nil {
	// 	return nil, err
	// }

	channel, err := conn.Channel()
	if err != nil {
		_ = conn.Close()
		return nil, err
	}

	return &RabbitMqConsumerImpl{
		conn:      conn,
		channel:   channel,
		queueName: queue,
		storage:   storage,
	}, nil
}

func (consumer *RabbitMqConsumerImpl) ConsumeMessages(handler func(message []byte)) error {
	_, err := consumer.channel.QueueDeclare(
		consumer.queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	messages, err := consumer.channel.Consume(
		consumer.queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range messages {
			handler(msg.Body)
		}
	}()

	return nil
}

func (consumer *RabbitMqConsumerImpl) Close() {
	if err := consumer.channel.Close(); err != nil {
		log.Printf("Failed to close RabbitMQ channel: %v", err)
	}
	if err := consumer.conn.Close(); err != nil {
		log.Printf("Failed to close RabbitMQ connection: %v", err)
	}
}
