package main

import (
	"context"

	unit_of_work "github.com/gabrielsc1998/go-ddd/internal/common/infra/db/unit-of-work"
	event_controller "github.com/gabrielsc1998/go-ddd/internal/events/infra/controllers/event"
	partner_controller "github.com/gabrielsc1998/go-ddd/internal/events/infra/controllers/partner"
	webserver "github.com/gabrielsc1998/go-ddd/internal/server"
)

func main() {
	db, err := SetupDatabase()
	panicIfHasError(err)

	uow := unit_of_work.NewUow(context.Background(), db.DB)
	RegisterRepositoriesInUOW(uow)

	services := SetupApplicationService(uow, db)

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
