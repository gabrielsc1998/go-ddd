package customer_repository

import entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/customer"

type CustomerRepositoryInterface interface {
	Add(customer *entity.Customer) error
	FindById(id string) (*entity.Customer, error)
	FindAll() ([]*entity.Customer, error)
	Delete(customer *entity.Customer) error
}
