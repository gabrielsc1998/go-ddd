package main

import (
	"fmt"
	"sync"

	application_service "github.com/gabrielsc1998/go-ddd/internal/common/application/application-service"
	event_dispatcher "github.com/gabrielsc1998/go-ddd/internal/common/application/events/event-dispatcher"
	event_handler "github.com/gabrielsc1998/go-ddd/internal/common/application/events/event-handler"
	domain_event "github.com/gabrielsc1998/go-ddd/internal/common/domain/domain-event"
	domain_event_manager "github.com/gabrielsc1998/go-ddd/internal/common/domain/domain-event-manager"
	unit_of_work "github.com/gabrielsc1998/go-ddd/internal/common/infra/db/unit-of-work"
	"github.com/gabrielsc1998/go-ddd/internal/database"
	event_service "github.com/gabrielsc1998/go-ddd/internal/events/application/services/event"
	partner_service "github.com/gabrielsc1998/go-ddd/internal/events/application/services/partner"
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
	"gorm.io/gorm"
)

func SetupDatabase() (*database.Database, error) {
	db := database.NewDatabase()
	err := db.ConnectMySQL(database.DatabaseMySQLOptions{
		Host:     "localhost",
		Port:     "3307",
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
	return db, err
}

func RegisterRepositoriesInUOW(uow *unit_of_work.Uow) {
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

type ApplicationServices struct {
	EventService   event_service.EventService
	PartnerService partner_service.PartnerService
}

func SetupApplicationService(uow *unit_of_work.Uow, db *database.Database) ApplicationServices {
	eventDispatcher := event_dispatcher.NewEventDispatcher()
	domainEventManager := domain_event_manager.NewDomainEventManager(eventDispatcher)
	applicationService := application_service.NewApplicationService(domainEventManager)

	domainEventManager.Register("PartnerCreated", event_handler.NewEventHandler(func(event domain_event.DomainEvent, wg *sync.WaitGroup) {
		fmt.Println("PartnerCreated")
		wg.Done()
	}))

	partnerRepository := partner_repository.NewPartnerRepository(db.DB)
	eventService := event_service.NewEventService(event_service.EventServiceProps{
		UOW:               uow,
		EventRepository:   event_repository.NewEventRepository(db.DB),
		PartnerRepository: partnerRepository,
	})
	partnerService := partner_service.NewPartnerService(partner_service.PartnerServiceProps{
		UOW:                uow,
		PartnerRepository:  partnerRepository,
		ApplicationService: applicationService,
	})
	return ApplicationServices{
		EventService:   eventService,
		PartnerService: partnerService,
	}
}
