package order_service

import (
	"errors"

	order_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/order"
)

type MockOrderRepo struct{}

func NewMockOrderRepo() *MockOrderRepo {
	return &MockOrderRepo{}
}

func (o *MockOrderRepo) Add(order *order_entity.Order) error {
	return errors.New("")
}

func (o *MockOrderRepo) FindById(id string) (*order_entity.Order, error) {
	return nil, errors.New("")
}

func (o *MockOrderRepo) FindAll(eventId string) ([]*order_entity.Order, error) {
	return nil, errors.New("")
}

func (o *MockOrderRepo) Delete(customer *order_entity.Order) error {
	return errors.New("")
}
