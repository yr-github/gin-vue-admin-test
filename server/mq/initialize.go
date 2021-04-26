package mq

import (
	"fmt"
	"gin-vue-admin/config"
	"go.uber.org/zap"

	"github.com/streadway/amqp"
)
type RabbitMQ struct {
	conn *amqp.Connection
	ch *amqp.Channel
	config config.Server
	log *zap.Logger
}

func Rabbit(config config.Server,log *zap.Logger)  *RabbitMQ {
	rabbit := &RabbitMQ{}
	rabbit.config = config
	rabbit.log = log
	rabbitConfig := rabbit.config.Rabbit
	url :=fmt.Sprintf("amqp://%s:%s@%s/",rabbitConfig.User,rabbitConfig.Password,rabbitConfig.Addr)
	var err error
	rabbit.conn, err = amqp.Dial(url)
	if err != nil {
		rabbit.log.Info(fmt.Sprintf("%s,%s", "Failed to connect to RabbitMQ", err))
	}
	//defer conn.Close()
	rabbit.ch, err = rabbit.conn.Channel()
	if err != nil {
		rabbit.log.Info(fmt.Sprintf("%s,%s", "Failed to open a channel", err))
	}
	//	defer ch.Close()
	return rabbit
}
