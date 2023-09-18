package main

import (
	"context"
	"encoding/json"

	"github.com/gabrielsc1998/go-ddd/cmd/setup"
	unit_of_work "github.com/gabrielsc1998/go-ddd/internal/common/infra/db/unit-of-work"
	"github.com/gabrielsc1998/go-ddd/internal/common/infra/outbox"
	customer_controller "github.com/gabrielsc1998/go-ddd/internal/events/infra/controllers/customer"
	event_controller "github.com/gabrielsc1998/go-ddd/internal/events/infra/controllers/event"
	partner_controller "github.com/gabrielsc1998/go-ddd/internal/events/infra/controllers/partner"
	webserver "github.com/gabrielsc1998/go-ddd/internal/server"
)

func main() {
	db, err := setup.SetupDatabase()
	panicIfHasError(err)

	rabbitmq := setup.SetupRabbitMq()
	ob := setup.SetupTransactionalOutbox(db, func(outboxData *[]outbox.OutboxModel, ob *outbox.Outbox) error {
		if outboxData == nil || len(*outboxData) == 0 {
			return nil
		}
		type Event struct {
			Name string `json:"name"`
		}
		currentEvent := Event{}
		err := json.Unmarshal([]byte((*outboxData)[0].Data), &currentEvent)
		if err != nil {
			return err
		}

		key := ""
		if currentEvent.Name == "PartnerCreatedInt" {
			key = "partner.created"
		} else if currentEvent.Name == "EventCreatedInt" {
			key = "event.created"
		}
		rabbitmq.Publish("events", key, (*outboxData)[0].Data)
		ob.MarkAsProcessed((*outboxData)[0].Id)
		return nil
	})

	uow := unit_of_work.NewUow(context.Background(), db.DB)
	setup.RegisterRepositoriesInUOW(uow)

	services := setup.SetupApplicationService(uow, db, ob)

	webserver := webserver.NewWebServer("8080")

	eventController := event_controller.NewEventController(services.EventService, services.OrderService)
	webserver.AddHandler("/events", "POST", eventController.CreateEvent)
	webserver.AddHandler("/events", "GET", eventController.ListEvents)
	webserver.AddHandler("/events/{event_id}/sections", "GET", eventController.FindEventSections)
	webserver.AddHandler("/events/{event_id}/sections", "POST", eventController.AddSection)
	webserver.AddHandler("/events/{event_id}/publish-all", "PUT", eventController.PublishAll)
	webserver.AddHandler("/events/{event_id}/sections/{section_id}", "PUT", eventController.UpdateSection)
	webserver.AddHandler("/events/{event_id}/sections/{section_id}/spots", "GET", eventController.GetSectionSpots)
	webserver.AddHandler("/events/{event_id}/sections/{section_id}/spots/{spot_id}", "PUT", eventController.UpdateLocation)
	webserver.AddHandler("/events/{event_id}/orders", "POST", eventController.CreateOrder)
	webserver.AddHandler("/events/{event_id}/orders", "GET", eventController.ListOrders)

	partnerController := partner_controller.NewPartnerController(services.PartnerService)
	webserver.AddHandler("/partners", "POST", partnerController.CreatePartner)
	webserver.AddHandler("/partners", "GET", partnerController.ListPartners)

	customerController := customer_controller.NewCustomerController(services.CustomerService)
	webserver.AddHandler("/customers", "POST", customerController.RegisterCustomer)
	webserver.AddHandler("/customers", "PUT", customerController.UpdateCustomer)
	webserver.AddHandler("/customers", "GET", customerController.FindCustomers)

	panicIfHasError(err)

	webserver.Start()
}

func panicIfHasError(err error) {
	if err != nil {
		panic(err)
	}
}
