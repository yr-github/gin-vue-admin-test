package mq

import (
	"fmt"
	"github.com/streadway/amqp"
)

func (rabbit *RabbitMQ)Send(msg string) error{
	err := rabbit.ch.Publish(
		"",     // exchange
		rabbit.queue.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	if err != nil {
		rabbit.log.Info(fmt.Sprintf("%s,%s", "Failed to publish a message", err))
	}
	rabbit.log.Info(fmt.Sprintf(" [x] Sent %s", msg))
	return err
}
