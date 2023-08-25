package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	Channel *amqp.Channel
}

type RabbitMQOptions struct {
	User     string
	Password string
	Host     string
	Port     string
}

func NewRabbitMQ() *RabbitMQ {
	return &RabbitMQ{}
}

func (r *RabbitMQ) Connect(options RabbitMQOptions) error {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", options.User, options.Password, options.Host, options.Port)
	conn, err := amqp.Dial(url)
	if err != nil {
		return err
	}
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	r.Channel = ch
	return nil
}
