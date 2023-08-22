package order_mapper

import (
	order_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/order"
	order_model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/order"
)

type OrderMapper struct {
}

func NewOrderMapper() *OrderMapper {
	return &OrderMapper{}
}

func (mapper *OrderMapper) ToModel(entity *order_entity.Order) *order_model.Order {
	return &order_model.Order{
		ID:          entity.Id.Value,
		CustomerId:  entity.CustomerId.Value,
		Amount:      entity.Amount,
		EventSpotId: entity.EventSpotId.Value,
		Status:      entity.Status,
	}
}

func (mapper *OrderMapper) ToEntity(model *order_model.Order) *order_entity.Order {
	section, _ := order_entity.Create(order_entity.OrderCreateProps{
		Id:          model.ID,
		CustomerId:  model.CustomerId,
		Amount:      model.Amount,
		EventSpotId: model.EventSpotId,
		Status:      model.Status,
	})
	return section
}
