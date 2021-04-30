package global

import (
	"fmt"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

type RabbitMQ struct {
	conn  *amqp.Connection
	ch    *amqp.Channel
	queue amqp.Queue
}

func Rabbit(queue string, durable bool, autoDelete bool) *RabbitMQ {
	rabbit := &RabbitMQ{}
	rabbitConfig := GVA_CONFIG.Rabbit
	url := fmt.Sprintf("amqp://%s:%s@%s/", rabbitConfig.User, rabbitConfig.Password, rabbitConfig.Addr)
	var err error
	rabbit.conn, err = amqp.Dial(url)
	if err != nil {
		GVA_LOG.Info(fmt.Sprintf("%s,%s", "Failed to connect to RabbitMQ", err))
	}
	//defer conn.Close()
	rabbit.ch, err = rabbit.conn.Channel()
	if err != nil {
		GVA_LOG.Info(fmt.Sprintf("%s,%s", "Failed to open a channel", err))
	}
	//	defer ch.Close()
	rabbit.queue, err = rabbit.ch.QueueDeclare(
		queue,      // name
		durable,    // durable
		autoDelete, // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		GVA_LOG.Info(fmt.Sprintf("%s,%s", "Failed to declare a queue", err))
	}
	return rabbit
}

func (rabbit *RabbitMQ) MqSend(msg string, queue string) error {
	err := rabbit.ch.Publish(
		"",    // exchange
		queue, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	if err != nil {
		GVA_LOG.Info(fmt.Sprintf("%s,%s", "Failed to publish a message", err))
	}
	GVA_LOG.Info(fmt.Sprintf(" [x] Sent %s", msg))
	return err
}

func (rabbit *RabbitMQ) MqReceive(queuename string) {
	go func() {
		msgs, err := rabbit.ch.Consume(
			queuename, // queue
			"",        // consumer
			false,     // auto-ack
			false,     // exclusive
			false,     // no-local
			false,     // no-wait
			nil,       // args
		)
		if err != nil {
			GVA_LOG.Info(fmt.Sprintf("%s,%s", "Failed to register a consumer", err))
		}
		for d := range msgs {
			GVA_LOG.Info(fmt.Sprintf("Received a message: %s", d.Body))
			if err != nil {
				GVA_LOG.Error("rabbit获取失败!", zap.Any("err", err))
			} else {
				err := d.Ack(false)
				if err != nil {
					//TODO error debug变为了逆序，未查到原因
					GVA_LOG.Error("rabbit ACK失败!", zap.Any("err", err))
				} else {
					//传值 这里传过去的值会被mqt2db自动处理找到可执行函数
					MQTODB <- string(d.Body)
				}
			}
		}
	}()
	//此处之所以案例代码要使用阻塞
	//是因为为了防止main退出，而我们的程序不需要。
	//<-make(chan bool)
}
