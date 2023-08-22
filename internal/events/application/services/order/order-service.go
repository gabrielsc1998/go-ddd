package order_service

import (
	"context"
	"errors"
	"time"

	unit_of_work "github.com/gabrielsc1998/go-ddd/internal/common/infra/db/unit-of-work"
	order_dto "github.com/gabrielsc1998/go-ddd/internal/events/application/dto/order"
	customer_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/customer"
	event_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/event"
	order_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/order"
	spot_reservation_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/spot-reservation"
	customer_repository "github.com/gabrielsc1998/go-ddd/internal/events/domain/repositories/customer"
	event_repository "github.com/gabrielsc1998/go-ddd/internal/events/domain/repositories/event"
	order_repository "github.com/gabrielsc1998/go-ddd/internal/events/domain/repositories/order"
	spot_reservation_repository "github.com/gabrielsc1998/go-ddd/internal/events/domain/repositories/spot-reservation"
)

type OrderService struct {
	uow                       *unit_of_work.Uow
	orderRepository           order_repository.OrderRepositoryInterface
	customerRepository        customer_repository.CustomerRepositoryInterface
	eventRepository           event_repository.EventRepositoryInterface
	spotReservationRepository spot_reservation_repository.SpotReservationRepositoryInterface
}

type OrderServiceProps struct {
	UOW                       *unit_of_work.Uow
	OrderRepository           order_repository.OrderRepositoryInterface
	CustomerRepository        customer_repository.CustomerRepositoryInterface
	EventRepository           event_repository.EventRepositoryInterface
	SpotReservationRepository spot_reservation_repository.SpotReservationRepositoryInterface
}

func NewOrderService(props OrderServiceProps) OrderService {
	return OrderService{
		uow:                       props.UOW,
		orderRepository:           props.OrderRepository,
		customerRepository:        props.CustomerRepository,
		eventRepository:           props.EventRepository,
		spotReservationRepository: props.SpotReservationRepository,
	}
}

func (o *OrderService) getOrderRepository() (order_repository.OrderRepositoryInterface, error) {
	ctx := context.Background()
	repo, err := o.uow.GetRepository(ctx, "OrderRepository")
	if err != nil {
		return nil, err
	}
	orderRepository := repo.(order_repository.OrderRepositoryInterface)
	return orderRepository, nil
}

func (o *OrderService) List() ([]*order_entity.Order, error) {
	orders, err := o.orderRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (o *OrderService) Create(input order_dto.OrderCreateDto) error {
	customer, err := o.getCustomer(input.CustomerId)
	if err != nil {
		return err
	}

	event, err := o.getEvent(input.EventId)
	if err != nil {
		return err
	}

	allowReserverSpot := event.AllowReserveSpot(event_entity.EventCommandReserveSpot{
		SectionId: input.SectionId,
		SpotId:    input.SpotId,
	})

	if !allowReserverSpot {
		return errors.New("spot not available")
	}

	spotReservationExists, _ := o.spotReservationRepository.FindById(input.SpotId)
	if spotReservationExists != nil {
		return errors.New("spot already reserved")
	}

	section, err := event.GetSection(input.SectionId)
	if err != nil {
		return err
	}

	order, err := order_entity.Create(order_entity.OrderCreateProps{
		CustomerId:  customer.Id.Value,
		EventSpotId: input.SpotId,
		Amount:      section.Price,
	})
	if order == nil {
		return err
	}

	order.Pay()

	errorInReservation := func() error {
		order.Cancel()
		err := o.orderRepository.Add(order)
		if err != nil {
			return err
		}
		return errors.New("error in reservation")
	}

	orderRepository, err := o.getOrderRepository()
	return o.uow.Do(o.uow.GetCtx(), func(uow *unit_of_work.Uow) error {
		err = orderRepository.Add(order)
		if err != nil {
			return errorInReservation()
		}

		spotReservation, err := spot_reservation_entity.Create(spot_reservation_entity.SpotReservationCreateProps{
			Id:              "",
			SpotId:          input.SpotId,
			CustomerId:      input.CustomerId,
			ReservationDate: time.Now(),
		})
		if err != nil {
			return errorInReservation()
		}
		err = o.spotReservationRepository.Add(spotReservation)
		if err != nil {
			return errorInReservation()
		}

		event.ReserveSpot(event_entity.EventCommandReserveSpot{
			SectionId: input.SectionId,
			SpotId:    input.SpotId,
		})
		err = o.eventRepository.Add(event)
		if err != nil {
			return errorInReservation()
		}
		return nil
	})
}

func (o *OrderService) getCustomer(customerId string) (*customer_entity.Customer, error) {
	customer, err := o.customerRepository.FindById(customerId)
	if err != nil {
		return nil, errors.New("customer not found")
	}
	return customer, nil
}

func (o *OrderService) getEvent(eventId string) (*event_entity.Event, error) {
	event, err := o.eventRepository.FindById(eventId)
	if err != nil {
		return nil, errors.New("event not found")
	}
	return event, nil
}
