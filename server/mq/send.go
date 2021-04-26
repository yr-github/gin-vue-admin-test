package mq

import (
	"fmt"
	"github.com/streadway/amqp"
)

func (rabbit *RabbitMQ)Send() {

	q, err := rabbit.ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err!=nil{
		rabbit.log.Info(fmt.Sprintf("%s,%s", "Failed to declare a queue", err))
	}
	body := "Hello World!"
	err = rabbit.ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		rabbit.log.Info(fmt.Sprintf("%s,%s", "Failed to publish a message", err))
	}
	rabbit.log.Info(fmt.Sprintf(" [x] Sent %s", body))
}
