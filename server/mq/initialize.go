package mq

import (
	"fmt"
	"gin-vue-admin/config"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/streadway/amqp"
)
type RabbitMQ struct {
	conn *amqp.Connection
	ch *amqp.Channel
	queue   amqp.Queue
	config config.Server
	log *zap.Logger
	db *gorm.DB
}

func Rabbit(config config.Server,log *zap.Logger,db *gorm.DB)  *RabbitMQ {
	rabbit := &RabbitMQ{}
	rabbit.config = config
	rabbit.log = log
	rabbit.db = db
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
	rabbit.queue, err = rabbit.ch.QueueDeclare(
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
	return rabbit
}
