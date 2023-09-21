package order_service

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/gabrielsc1998/go-ddd/internal/common/tests"
	order_dto "github.com/gabrielsc1998/go-ddd/internal/events/application/dto/order"
	customer_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/customer"
	event_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/event"
	order_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/order"
	partner_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/partner"
	section_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/section"
	spot_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/spot"
	customer_repository "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/repositories/customer"
	event_repository "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/repositories/event"
	order_repository "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/repositories/order"
	partner_repository "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/repositories/partner"
	spot_reservation_repository "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/repositories/spot-reservation"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	test            *tests.Tests
	orderService    OrderService
	event           event_entity.Event
	customer        customer_entity.Customer
	section         section_entity.Section
	spot            spot_entity.Spot
	orderRepository *order_repository.OrderRepository
)

func Setup() {
	test = tests.Setup()
	test.UOW.Register("OrderRepository", func(db *gorm.DB) interface{} {
		repo := order_repository.NewOrderRepository(db)
		return repo
	})
	test.UOW.Register("EventRepository", func(db *gorm.DB) interface{} {
		repo := event_repository.NewEventRepository(db)
		return repo
	})
	test.UOW.Register("SpotReservationRepository", func(db *gorm.DB) interface{} {
		repo := spot_reservation_repository.NewSpotReservationRepository(db)
		return repo
	})
	orderRepository = order_repository.NewOrderRepository(test.DB.DB)
	partnerRepository := partner_repository.NewPartnerRepository(test.DB.DB)
	customerRepository := customer_repository.NewCustomerRepository(test.DB.DB)
	eventRepository := event_repository.NewEventRepository(test.DB.DB)
	spotReservationRepository := spot_reservation_repository.NewSpotReservationRepository(test.DB.DB)
	orderService = NewOrderService(OrderServiceProps{
		UOW:                       test.UOW,
		OrderRepository:           orderRepository,
		CustomerRepository:        customerRepository,
		EventRepository:           eventRepository,
		SpotReservationRepository: spotReservationRepository,
	})

	createdPartner, _ := partner_entity.Create(partner_entity.PartnerCreateProps{
		Id:   "",
		Name: "Test Partner",
	})
	partnerRepository.Add(createdPartner)

	createdEvent, _ := event_entity.Create(event_entity.EventCreateProps{
		Id:          "",
		Name:        "Test Event",
		Description: "Test Event Description",
		Date:        time.Now(),
		IsPublished: true,
		PartnerId:   createdPartner.Id.Value,
	})

	createdEvent.AddSection(section_entity.SectionCreateProps{
		Id:                 "",
		Name:               "Test Section",
		Description:        "Test Section Description",
		Price:              100,
		TotalSpots:         10,
		TotalSpotsReserved: 0,
		IsPublished:        true,
		Date:               time.Now(),
	})
	createdEvent.PublishAll()
	err := eventRepository.Add(createdEvent)
	fmt.Println("erro aqui", err)
	event = *createdEvent

	section = event.Sections[0]
	spot = section.Spots[0]

	createdCustomer, _ := customer_entity.Create(customer_entity.CustomerCreateProps{
		Id:   "",
		Name: "Jhon Doe",
		CPF:  "45616278041",
	})
	customerRepository.Add(createdCustomer)
	customer = *createdCustomer
}

func CreateOrder() error {
	err := orderService.Create(order_dto.OrderCreateDto{
		EventId:    event.Id.Value,
		CustomerId: customer.Id.Value,
		SectionId:  section.Id.Value,
		SpotId:     spot.Id.Value,
		CardToken:  "tok_visa",
	})
	return err
}

func TestShouldReturnAnErrorWhenTheCustomerDoesNotExists(t *testing.T) {
	Setup()
	customer.Id.Value = "5225dd38-1a8d-429c-9eca-daf62caf1efd"
	err := CreateOrder()
	assert.EqualError(t, errors.New("customer not found"), err.Error())
}

func TestShouldReturnAnErrorWhenTheEventDoesNotExists(t *testing.T) {
	Setup()
	event.Id.Value = "5225dd38-1a8d-429c-9eca-daf62caf1efd"
	err := CreateOrder()
	assert.EqualError(t, errors.New("event not found"), err.Error())
}

func TestShouldCreateAnOrder(t *testing.T) {
	Setup()

	err := CreateOrder()
	assert.Nil(t, err)

	orders, _ := orderRepository.FindAll(event.Id.Value)
	assert.Equal(t, 1, len(orders))
	assert.Equal(t, spot.Id.Value, orders[0].EventSpotId.Value)
	assert.Equal(t, customer.Id.Value, orders[0].CustomerId.Value)
	assert.Equal(t, order_entity.OrderStatus.PAID, orders[0].Status)
	assert.Equal(t, float64(100), orders[0].Amount)
}

func TestShouldReturnErrorWhenTheSpotIsAlreadyReserved(t *testing.T) {
	Setup()

	err := CreateOrder()
	assert.Nil(t, err)

	err = CreateOrder()
	assert.EqualError(t, errors.New("spot not available"), err.Error())
}

func TestShouldCreateACanceledOrder(t *testing.T) {
	Setup()

	test.UOW.Register("OrderRepository", func(db *gorm.DB) interface{} {
		repo := NewMockOrderRepo()
		return repo
	})

	err := CreateOrder()
	assert.EqualError(t, errors.New("error in reservation"), err.Error())

	orders, _ := orderRepository.FindAll(event.Id.Value)
	assert.Equal(t, 1, len(orders))
	assert.Equal(t, order_entity.OrderStatus.CANCELED, orders[0].Status)
}

func TestShouldListOrders(t *testing.T) {
	Setup()

	err := CreateOrder()
	assert.Nil(t, err)

	orders, _ := orderRepository.FindAll(event.Id.Value)
	assert.Equal(t, 1, len(orders))
	assert.Equal(t, spot.Id.Value, orders[0].EventSpotId.Value)
	assert.Equal(t, customer.Id.Value, orders[0].CustomerId.Value)
	assert.Equal(t, order_entity.OrderStatus.PAID, orders[0].Status)
	assert.Equal(t, float64(100), orders[0].Amount)

	// orders, err = orderService.List("3e320fcd-e668-4cc5-8650-c33caea2bd56")
	// assert.Nil(t, err)
	// assert.Equal(t, 0, len(orders))
}
