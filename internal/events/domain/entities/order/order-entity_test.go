package order_entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCreateAnOrder(t *testing.T) {
	order, err := Create(OrderCreateProps{
		Id:          "",
		CustomerId:  "f2fc2af8-b2cb-418a-bb48-db1d5f8ed723",
		Amount:      0,
		EventId:     "4ea6b3c6-36d3-407a-be40-2eedcf590812",
		EventSpotId: "5810c547-b679-4f88-8258-90f01be933f2",
	})

	assert.Nil(t, err)
	assert.NotNil(t, order)
	assert.NotNil(t, order.Id.Value)
	assert.Equal(t, "f2fc2af8-b2cb-418a-bb48-db1d5f8ed723", order.CustomerId.Value)
	assert.Equal(t, "4ea6b3c6-36d3-407a-be40-2eedcf590812", order.EventId.Value)
	assert.Equal(t, "5810c547-b679-4f88-8258-90f01be933f2", order.EventSpotId.Value)
	assert.Equal(t, OrderStatus.PENDING, order.Status)
	assert.Equal(t, 0.0, order.Amount)
}

func TestShouldPayOrder(t *testing.T) {
	order, err := Create(OrderCreateProps{
		Id:          "",
		CustomerId:  "f2fc2af8-b2cb-418a-bb48-db1d5f8ed723",
		Amount:      0,
		EventId:     "4ea6b3c6-36d3-407a-be40-2eedcf590812",
		EventSpotId: "5810c547-b679-4f88-8258-90f01be933f2",
	})
	assert.Nil(t, err)
	assert.Equal(t, OrderStatus.PENDING, order.Status)

	order.Pay()
	assert.Equal(t, OrderStatus.PAID, order.Status)
}

func TestShouldCancelOrder(t *testing.T) {
	order, err := Create(OrderCreateProps{
		Id:          "",
		CustomerId:  "f2fc2af8-b2cb-418a-bb48-db1d5f8ed723",
		Amount:      0,
		EventId:     "4ea6b3c6-36d3-407a-be40-2eedcf590812",
		EventSpotId: "5810c547-b679-4f88-8258-90f01be933f2",
	})
	assert.Nil(t, err)
	assert.Equal(t, OrderStatus.PENDING, order.Status)

	order.Cancel()
	assert.Equal(t, OrderStatus.CANCELED, order.Status)
}
