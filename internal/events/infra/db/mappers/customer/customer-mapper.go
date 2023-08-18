package customer_mapper

import (
	customer_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/customer"
	customer_model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/customer"
)

type CustomerMapper struct {
}

func NewCustomerMapper() *CustomerMapper {
	return &CustomerMapper{}
}

func (mapper *CustomerMapper) ToModel(entity *customer_entity.Customer) *customer_model.Customer {
	return &customer_model.Customer{
		ID:   entity.Id.Value,
		Name: entity.Name,
		CPF:  entity.CPF.Value,
	}
}

func (mapper *CustomerMapper) ToEntity(model *customer_model.Customer) *customer_entity.Customer {
	customer, _ := customer_entity.Create(customer_entity.CustomerCreateProps{
		Id:   model.ID,
		Name: model.Name,
		CPF:  model.CPF,
	})
	return customer
}
