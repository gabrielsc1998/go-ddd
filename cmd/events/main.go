package main

import (
	"context"

	"github.com/gabrielsc1998/go-ddd/cmd/setup"
	unit_of_work "github.com/gabrielsc1998/go-ddd/internal/common/infra/db/unit-of-work"
	"github.com/gabrielsc1998/go-ddd/internal/common/infra/outbox"
	event_controller "github.com/gabrielsc1998/go-ddd/internal/events/infra/controllers/event"
	partner_controller "github.com/gabrielsc1998/go-ddd/internal/events/infra/controllers/partner"
	webserver "github.com/gabrielsc1998/go-ddd/internal/server"
	"github.com/streadway/amqp"
)

func main() {
	db, err := setup.SetupDatabase()
	panicIfHasError(err)

	rabbitmq := setup.SetupRabbitMq()
	ob := setup.SetupTransactionalOutbox(db, func(outboxData *[]outbox.OutboxModel, ob *outbox.Outbox) error {
		if outboxData == nil || len(*outboxData) == 0 {
			return nil
		}
		rabbitmq.Channel.Publish("events", "partner.created", false, false, amqp.Publishing{
			ContentType: "application/json",
			Body:        (*outboxData)[0].Data,
		})
		ob.MarkAsProcessed((*outboxData)[0].Id)
		return nil
	})

	uow := unit_of_work.NewUow(context.Background(), db.DB)
	setup.RegisterRepositoriesInUOW(uow)

	services := setup.SetupApplicationService(uow, db, ob)

	webserver := webserver.NewWebServer("8080")

	eventController := event_controller.NewEventController(services.EventService)
	webserver.AddHandler("/events", "POST", eventController.CreateEvent)
	webserver.AddHandler("/events", "GET", eventController.ListEvents)

	partnerController := partner_controller.NewPartnerController(services.PartnerService)
	webserver.AddHandler("/partners", "POST", partnerController.CreatePartner)
	webserver.AddHandler("/partners", "GET", partnerController.ListPartners)

	panicIfHasError(err)

	webserver.Start()
}

func panicIfHasError(err error) {
	if err != nil {
		panic(err)
	}
}
