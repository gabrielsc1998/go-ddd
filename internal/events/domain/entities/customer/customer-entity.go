package customer_entity

import (
	"errors"

	"github.com/gabrielsc1998/go-ddd/internal/common/domain/entity"
	"github.com/gabrielsc1998/go-ddd/internal/common/domain/value-objects/cpf"
	"github.com/gabrielsc1998/go-ddd/internal/common/domain/value-objects/id"
)

type Customer struct {
	entity.Entity
	Name string   `json:"name"`
	CPF  *cpf.CPF `json:"cpf"`
}

type CustomerCreateProps struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	CPF  string `json:"cpf"`
}

func Create(props CustomerCreateProps) (*Customer, error) {
	err := validate(props)
	if err != nil {
		return nil, err
	}
	customerId, _ := id.NewID(props.Id)
	customerCPF, _ := cpf.NewCPF(props.CPF)
	return &Customer{
		Entity: entity.Entity{
			Id: customerId,
		},
		Name: props.Name,
		CPF:  customerCPF,
	}, nil
}

func validate(props CustomerCreateProps) error {
	_, err := id.NewID(props.Id)
	if err != nil {
		return err
	}
	if props.Name == "" {
		return errors.New("invalid name")
	}
	_, err = cpf.NewCPF(props.CPF)
	if err != nil {
		return err
	}
	return nil
}

func (c *Customer) ChangeCPF(newCPF string) error {
	err := validate(CustomerCreateProps{
		Id:   c.Id.Value,
		Name: c.Name,
		CPF:  newCPF,
	})
	if err != nil {
		return err
	}
	c.CPF, _ = cpf.NewCPF(newCPF)
	return nil
}

func (c *Customer) ChangeName(newName string) error {
	err := validate(CustomerCreateProps{
		Id:   c.Id.Value,
		Name: newName,
		CPF:  c.CPF.Value,
	})
	if err != nil {
		return err
	}
	c.Name = newName
	return nil
}
