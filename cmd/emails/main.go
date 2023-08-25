package main

import (
	"fmt"

	"github.com/gabrielsc1998/go-ddd/cmd/setup"
	"github.com/gabrielsc1998/go-ddd/internal/common/infra/rabbitmq"
)

func Consumer(rabbitmq *rabbitmq.RabbitMQ) {
	msgs, _ := rabbitmq.Channel.Consume("partner-created-queue", "", false, false, false, false, nil)
	for msg := range msgs {
		msg.Ack(false)
		fmt.Printf("Consumer partner-created received: %s", string(msg.Body))
	}
}

func main() {
	rabbitmq := setup.SetupRabbitMq()
	fmt.Println("Connected to RabbitMQ")
	Consumer(rabbitmq)
}
