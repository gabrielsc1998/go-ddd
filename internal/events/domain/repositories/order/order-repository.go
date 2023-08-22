package order_repository

import entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/order"

type OrderRepositoryInterface interface {
	Add(order *entity.Order) error
	FindById(id string) (*entity.Order, error)
	FindAll() ([]*entity.Order, error)
	Delete(order *entity.Order) error
}
