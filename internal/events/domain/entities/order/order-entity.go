package order_entity

import (
	"github.com/gabrielsc1998/go-ddd/internal/common/domain/entity"
	"github.com/gabrielsc1998/go-ddd/internal/common/domain/value-objects/id"
)

var OrderStatus = struct {
	PENDING  int
	PAID     int
	CANCELED int
}{
	PENDING:  0,
	PAID:     1,
	CANCELED: 2,
}

type Order struct {
	entity.AggregateRoot
	CustomerId  *id.ID
	Amount      float64
	EventSpotId *id.ID
	Status      int
}

type OrderCreateProps struct {
	Id          string
	CustomerId  string
	Amount      float64
	EventSpotId string
	Status      int
}

func Create(props OrderCreateProps) (*Order, error) {
	err := validate(props)
	if err != nil {
		return nil, err
	}
	orderId, _ := id.NewID(props.Id)
	customerId, _ := id.NewID(props.CustomerId)
	eventSpotId, _ := id.NewID(props.EventSpotId)
	status := OrderStatus.PENDING
	if props.Status != OrderStatus.PENDING {
		status = props.Status
	}
	return &Order{
		AggregateRoot: entity.AggregateRoot{
			Id: orderId,
		},
		CustomerId:  customerId,
		Amount:      props.Amount,
		EventSpotId: eventSpotId,
		Status:      status,
	}, nil
}

func validate(props OrderCreateProps) error {
	_, err := id.NewID(props.Id)
	if err != nil {
		return err
	}
	_, err = id.NewID(props.CustomerId)
	if err != nil {
		return err
	}
	_, err = id.NewID(props.EventSpotId)
	if err != nil {
		return err
	}
	return nil
}

func (o *Order) Pay() {
	o.Status = OrderStatus.PAID
}

func (o *Order) Cancel() {
	o.Status = OrderStatus.CANCELED
}
