package mq

import "fmt"

func (rabbit *RabbitMQ) Receive()  {
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
	msgs, err := rabbit.ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err!=nil{
		rabbit.log.Info(fmt.Sprintf("%s,%s", "Failed to register a consumer", err))
	}
	// create an unbuffered channel for bool types.
	// Type is not important but we have to give one anyway.
	forever := make(chan bool)
	// fire up a goroutine that hooks onto msgs channel and reads
	// anything that pops into it. This essentially is a thread of
	// execution within the main thread. msgs is a channel constructed by
	// previous code.
	go func() {
		for d := range msgs {
			rabbit.log.Info(fmt.Sprintf("Received a message: %s", d.Body))
		}
	}()
	if err!=nil{
		rabbit.log.Info(fmt.Sprintf(" [*] Waiting for messages. To exit press CTRL+C", err))
	}
	// We need to block the main thread so that the above thread stays
	// on reading from msgs channel. To do that just try to read in from
	// the forever channel. As long as no one writes to it we will wait here.
	// Since we are the only ones that know of it it is guaranteed that
	// nothing gets written in it. We could also do a busy wait here but
	// that would waste CPU cycles for no good reason.
	<-forever
}
