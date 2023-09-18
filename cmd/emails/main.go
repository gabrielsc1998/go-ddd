package main

import (
	"fmt"

	"github.com/gabrielsc1998/go-ddd/cmd/setup"
	"github.com/gabrielsc1998/go-ddd/internal/common/infra/rabbitmq"
)

func Consumer(rabbitmq *rabbitmq.RabbitMQ, queue string) {
	msgs, _ := rabbitmq.Consume(queue)
	for msg := range msgs {
		msg.Ack(false)
		fmt.Printf("Consumer %s received: %s \n\n", queue, string(msg.Body))
	}
}

func main() {
	rabbitmq := setup.SetupRabbitMq()
	go Consumer(rabbitmq, "event-created-queue")
	go Consumer(rabbitmq, "partner-created-queue")
	for {
	}
}
