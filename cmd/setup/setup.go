package setup

import (
	"encoding/json"
	"fmt"
	"sync"

	application_service "github.com/gabrielsc1998/go-ddd/internal/common/application/application-service"
	event_handler "github.com/gabrielsc1998/go-ddd/internal/common/application/events/event-handler"
	domain_event "github.com/gabrielsc1998/go-ddd/internal/common/domain/domain-event"
	domain_event_manager "github.com/gabrielsc1998/go-ddd/internal/common/domain/domain-event-manager"
	unit_of_work "github.com/gabrielsc1998/go-ddd/internal/common/infra/db/unit-of-work"
	"github.com/gabrielsc1998/go-ddd/internal/common/infra/outbox"
	"github.com/gabrielsc1998/go-ddd/internal/common/infra/rabbitmq"
	"github.com/gabrielsc1998/go-ddd/internal/database"
	customer_service "github.com/gabrielsc1998/go-ddd/internal/events/application/services/customer"
	event_service "github.com/gabrielsc1998/go-ddd/internal/events/application/services/event"
	partner_service "github.com/gabrielsc1998/go-ddd/internal/events/application/services/partner"
	event_int_events "github.com/gabrielsc1998/go-ddd/internal/events/domain/events/event/integration-events"
	partner_int_events "github.com/gabrielsc1998/go-ddd/internal/events/domain/events/partner/integration-events"
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

func panicIfHasError(err error) {
	if err != nil {
		panic(err)
	}
}

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
	EventService    event_service.EventService
	PartnerService  partner_service.PartnerService
	CustomerService customer_service.CustomerService
}

func SetupRabbitMq() *rabbitmq.RabbitMQ {
	rmq := rabbitmq.NewRabbitMQ()
	err := rmq.Connect(rabbitmq.RabbitMQOptions{
		User:     "guest",
		Password: "guest",
		Host:     "localhost",
		Port:     "5672",
	})
	panicIfHasError(err)
	fmt.Println("Connected to RabbitMQ")
	rmq.Channel.ExchangeDeclare("events", "topic", true, false, false, false, nil)
	rmq.Channel.QueueDeclare("partner-created-queue", true, false, false, false, nil)
	rmq.Channel.QueueBind("partner-created-queue", "partner.created", "events", false, nil)
	rmq.Channel.QueueDeclare("event-created-queue", true, false, false, false, nil)
	rmq.Channel.QueueBind("event-created-queue", "event.created", "events", false, nil)
	return rmq
}

func SetupTransactionalOutbox(db *database.Database, handle func(outboxData *[]outbox.OutboxModel, ob *outbox.Outbox) error) *outbox.Outbox {
	return outbox.NewOutbox(db.DB, handle)
}

func SetupApplicationService(uow *unit_of_work.Uow, db *database.Database, ob *outbox.Outbox) ApplicationServices {
	domainEventManager := domain_event_manager.NewDomainEventManager()
	applicationService := application_service.NewApplicationService(domainEventManager)

	domainEventManager.RegisterForDomainEvent("PartnerCreated", event_handler.NewEventHandler(func(event interface{}, wg *sync.WaitGroup) {
		fmt.Println("PartnerCreated - domain")
		wg.Done()
	}))

	domainEventManager.RegisterForIntegrationEvent("PartnerCreated", event_handler.NewEventHandler(func(event interface{}, wg *sync.WaitGroup) {
		partnerCreatedIntEvent := partner_int_events.NewPartnerCreatedEvent(event.(*domain_event.DomainEvent))
		fmt.Println("PartnerCreated - integration", partnerCreatedIntEvent)
		payload, err := json.Marshal(partnerCreatedIntEvent)
		panicIfHasError(err)
		err = ob.Add(outbox.DtoAddInOutbox{
			Payload: payload,
		})
		panicIfHasError(err)
		wg.Done()
	}))

	domainEventManager.RegisterForDomainEvent("EventCreated", event_handler.NewEventHandler(func(event interface{}, wg *sync.WaitGroup) {
		fmt.Println("EventCreated - domain")
		wg.Done()
	}))

	domainEventManager.RegisterForIntegrationEvent("EventCreated", event_handler.NewEventHandler(func(event interface{}, wg *sync.WaitGroup) {
		eventCreatedIntEvent := event_int_events.NewEventCreatedEvent(event.(*domain_event.DomainEvent))
		fmt.Println("EventCreated - integration", eventCreatedIntEvent)
		payload, err := json.Marshal(eventCreatedIntEvent)
		panicIfHasError(err)
		err = ob.Add(outbox.DtoAddInOutbox{
			Payload: payload,
		})
		panicIfHasError(err)
		wg.Done()
	}))

	partnerRepository := partner_repository.NewPartnerRepository(db.DB)
	eventService := event_service.NewEventService(event_service.EventServiceProps{
		UOW:                uow,
		EventRepository:    event_repository.NewEventRepository(db.DB),
		PartnerRepository:  partnerRepository,
		ApplicationService: applicationService,
	})
	partnerService := partner_service.NewPartnerService(partner_service.PartnerServiceProps{
		UOW:                uow,
		PartnerRepository:  partnerRepository,
		ApplicationService: applicationService,
	})
	customerService := customer_service.NewCustomerService(customer_service.CustomerServiceProps{
		UOW:                uow,
		CustomerRepository: customer_repository.NewCustomerRepository(db.DB),
		// ApplicationService: applicationService,
	})

	return ApplicationServices{
		EventService:    eventService,
		PartnerService:  partnerService,
		CustomerService: customerService,
	}
}
