package main

import (
	"context"
	"fmt"

	unit_of_work "github.com/gabrielsc1998/go-ddd/internal/common/infra/db/unit-of-work"
	"github.com/gabrielsc1998/go-ddd/internal/database"
	event_service "github.com/gabrielsc1998/go-ddd/internal/events/application/services/event"
	partner_service "github.com/gabrielsc1998/go-ddd/internal/events/application/services/partner"
	event_controller "github.com/gabrielsc1998/go-ddd/internal/events/infra/controllers/event"
	partner_controller "github.com/gabrielsc1998/go-ddd/internal/events/infra/controllers/partner"
	customer_model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/customer"
	event_model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/event"
	order_model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/order"
	partner_model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/partner"
	section_model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/section"
	spot_model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/spot"
	spot_reservation_model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/spot-reservation"
	customer_repository "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/repositories/customer"
	event_repository "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/repositories/event"
	order_repository "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/repositories/order"
	partner_repository "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/repositories/partner"
	spot_reservation_repository "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/repositories/spot-reservation"
	webserver "github.com/gabrielsc1998/go-ddd/internal/server"
	"gorm.io/gorm"
)

func registerRepositoriesInUOW(uow *unit_of_work.Uow) {
	uow.Register("EventRepository", func(db *gorm.DB) interface{} {
		repo := event_repository.NewEventRepository(db)
		return repo
	})
	uow.Register("PartnerRepository", func(db *gorm.DB) interface{} {
		repo := partner_repository.NewPartnerRepository(db)
		return repo
	})
	uow.Register("CustomerRepository", func(db *gorm.DB) interface{} {
		repo := customer_repository.NewCustomerRepository(db)
		return repo
	})
	uow.Register("OrderRepository", func(db *gorm.DB) interface{} {
		repo := order_repository.NewOrderRepository(db)
		return repo
	})
	uow.Register("SpotReservationRepository", func(db *gorm.DB) interface{} {
		repo := spot_reservation_repository.NewSpotReservationRepository(db)
		return repo
	})
}

func main() {
	db := database.NewDatabase()
	err := db.ConnectMySQL(database.DatabaseMySQLOptions{
		Host:     "localhost",
		Port:     "3306",
		User:     "root",
		Password: "root",
		Database: "events",
	})
	panicIfHasError(err)
	fmt.Println("Connected to MySQL")
	err = db.DB.AutoMigrate(
		&event_model.Event{},
		&section_model.Section{},
		&spot_model.Spot{},
		&partner_model.Partner{},
		&customer_model.Customer{},
		&spot_reservation_model.SpotReservation{},
		&order_model.Order{},
	)
	panicIfHasError(err)

	uow := unit_of_work.NewUow(context.Background(), db.DB)
	registerRepositoriesInUOW(uow)

	eventService := event_service.NewEventService(event_service.EventServiceProps{
		UOW:               uow,
		EventRepository:   event_repository.NewEventRepository(db.DB),
		PartnerRepository: partner_repository.NewPartnerRepository(db.DB),
	})
	eventController := event_controller.NewEventController(eventService)

	partnerService := partner_service.NewPartnerService(partner_service.PartnerServiceProps{
		UOW:               uow,
		PartnerRepository: partner_repository.NewPartnerRepository(db.DB),
	})
	partnerController := partner_controller.NewPartnerController(partnerService)

	webserver := webserver.NewWebServer("8080")
	webserver.AddHandler("/events", "POST", eventController.CreateEvent)
	webserver.AddHandler("/events", "GET", eventController.ListEvents)
	webserver.AddHandler("/partners", "POST", partnerController.CreatePartner)
	webserver.AddHandler("/partners", "GET", partnerController.ListPartners)
	webserver.Start()
}

func panicIfHasError(err error) {
	if err != nil {
		panic(err)
	}
}
