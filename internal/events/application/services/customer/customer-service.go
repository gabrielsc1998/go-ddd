package customer_service

import (
	"context"

	unit_of_work "github.com/gabrielsc1998/go-ddd/internal/common/infra/db/unit-of-work"
	customer_dto "github.com/gabrielsc1998/go-ddd/internal/events/application/dto/customer"
	customer_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/customer"
	customer_repository "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/repositories/customer"
)

type CustomerService struct {
	uow                *unit_of_work.Uow
	customerRepository *customer_repository.CustomerRepository
}

type CustomerServiceProps struct {
	UOW                *unit_of_work.Uow
	CustomerRepository *customer_repository.CustomerRepository
}

func NewCustomerService(props CustomerServiceProps) CustomerService {
	return CustomerService{
		uow:                props.UOW,
		customerRepository: props.CustomerRepository,
	}
}

func (c *CustomerService) getCustomerRepository() (*customer_repository.CustomerRepository, error) {
	ctx := context.Background()
	repo, err := c.uow.GetRepository(ctx, "CustomerRepository")
	if err != nil {
		return nil, err
	}
	customerRepository := repo.(*customer_repository.CustomerRepository)
	return customerRepository, nil
}

func (c *CustomerService) Register(input customer_dto.CustomerRegisterDto) error {
	customer, _ := customer_entity.Create(customer_entity.CustomerCreateProps{
		Id:   "",
		Name: input.Name,
		CPF:  input.CPF,
	})
	customerRepository, err := c.getCustomerRepository()
	err = c.uow.Do(c.uow.GetCtx(), func(uow *unit_of_work.Uow) error {
		err = customerRepository.Add(customer)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (c *CustomerService) Update(input customer_dto.CustomerUpdateDto) error {
	customerRepository, err := c.getCustomerRepository()
	if err != nil {
		return err
	}
	customer, err := customerRepository.FindById(input.Id)
	if err != nil {
		return err
	}
	if input.Name != "" {
		customer.ChangeName(input.Name)
	}
	err = c.uow.Do(c.uow.GetCtx(), func(uow *unit_of_work.Uow) error {
		err = customerRepository.Add(customer)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (c *CustomerService) FindCustomers() ([]*customer_entity.Customer, error) {
	customers, err := c.customerRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return customers, nil
}
