package order_presenter

import order_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/order"

type OrderPresenter struct {
	Id          string  `json:"id"`
	CustomerId  string  `json:"customer_id"`
	Amount      float64 `json:"amount"`
	EventSpotId string  `json:"event_spot_id"`
	Status      int     `json:"status"`
}

func ToPresent(order *order_entity.Order) OrderPresenter {
	return OrderPresenter{
		Id:          order.Id.Value,
		CustomerId:  order.CustomerId.Value,
		Amount:      order.Amount,
		EventSpotId: order.EventSpotId.Value,
		Status:      order.Status,
	}
}
