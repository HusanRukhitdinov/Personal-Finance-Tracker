package rabbitMq

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
	_ "github.com/streadway/amqp"
)

type RabbitMqProduce interface {
	ProduceMassage(queueName string, message []byte) error
	Close()
}

type RabbitMqProducerInt struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewRabbitMqProducerInt(url string) (*RabbitMqProducerInt, error) {
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
	channel, _ := conn.Channel()
	return &RabbitMqProducerInt{
		conn:    conn,
		channel: channel,
	}, nil
}
func (rabbitMq *RabbitMqProducerInt) ProduceMassage(queueName string, message []byte) error {
	declareQueue, err := rabbitMq.channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		return err
	}

	err = rabbitMq.channel.Publish(
		"",
		declareQueue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
	if err != nil {
		return err
	}
	fmt.Printf(" [x] Sent", string(message))
	return nil
}
