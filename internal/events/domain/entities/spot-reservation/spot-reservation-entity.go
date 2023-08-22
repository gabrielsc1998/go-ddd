package spot_reservation_entity

import (
	"time"

	"github.com/gabrielsc1998/go-ddd/internal/common/domain/entity"
	"github.com/gabrielsc1998/go-ddd/internal/common/domain/value-objects/id"
)

type SpotReservation struct {
	entity.AggregateRoot
	SpotId          *id.ID
	CustomerId      *id.ID
	ReservationDate time.Time
}

type SpotReservationCreateProps struct {
	Id              string
	SpotId          string
	CustomerId      string
	ReservationDate time.Time
}

type SpotReservationCommandChangeReservation struct {
	CustomerId      string
	ReservationDate time.Time
}

func Create(props SpotReservationCreateProps) (*SpotReservation, error) {
	err := validate(props)
	if err != nil {
		return nil, err
	}
	orderId, _ := id.NewID(props.Id)
	customerId, _ := id.NewID(props.CustomerId)
	spotId, _ := id.NewID(props.SpotId)
	return &SpotReservation{
		AggregateRoot: entity.AggregateRoot{
			Id: orderId,
		},
		CustomerId:      customerId,
		SpotId:          spotId,
		ReservationDate: props.ReservationDate,
	}, nil
}

func validate(props SpotReservationCreateProps) error {
	_, err := id.NewID(props.Id)
	if err != nil {
		return err
	}
	_, err = id.NewID(props.CustomerId)
	if err != nil {
		return err
	}
	return nil
}

func (order *SpotReservation) ChangeReservation(command SpotReservationCommandChangeReservation) {
	order.CustomerId, _ = id.NewID(command.CustomerId)
	order.ReservationDate = command.ReservationDate
}
