package mq

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
)

func (rabbit *RabbitMQ) Receive()  {

	msgs, err := rabbit.ch.Consume(
		rabbit.queue.Name, // queue
		"",     // consumer
		false,   // auto-ack
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

	// fire up a goroutine that hooks onto msgs channel and reads
	// anything that pops into it. This essentially is a thread of
	// execution within the main thread. msgs is a channel constructed by
	// previous code.
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			rabbit.log.Info(fmt.Sprintf("Received a message: %s", d.Body))
			var mytask MyTask
			json.Unmarshal([]byte(d.Body),&mytask)

			if err := rabbit.db.Create(&mytask).Error; err != nil {
				rabbit.log.Error("创建失败!", zap.Any("err", err))
				//response.FailWithMessage("创建失败", c)
			} else {
				//response.OkWithMessage("创建成功", c)
				d.Ack(false)
			}

		}
		<-forever
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

}
