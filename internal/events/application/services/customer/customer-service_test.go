package customer_service

import (
	"testing"

	"github.com/gabrielsc1998/go-ddd/internal/common/tests"
	customer_dto "github.com/gabrielsc1998/go-ddd/internal/events/application/dto/customer"
	customer_repository "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/repositories/customer"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var test *tests.Tests
var customerService CustomerService

func Setup() {
	test = tests.Setup()
	test.UOW.Register("CustomerRepository", func(db *gorm.DB) interface{} {
		repo := customer_repository.NewCustomerRepository(db)
		return repo
	})
	customerService = NewCustomerService(CustomerServiceProps{
		UOW:                test.UOW,
		CustomerRepository: customer_repository.NewCustomerRepository(test.DB.DB),
	})
}

func TestShouldRegisterACustomer(t *testing.T) {
	Setup()

	err := customerService.Register(customer_dto.CustomerRegisterDto{
		Name: "Jhon Doe",
		CPF:  "45616278041",
	})
	assert.Nil(t, err)

	repo := customer_repository.NewCustomerRepository(test.DB.DB)
	customer, err := repo.FindAll()

	assert.Nil(t, err)
	assert.Equal(t, 1, len(customer))
	assert.Equal(t, "Jhon Doe", customer[0].Name)
	assert.Equal(t, "45616278041", customer[0].CPF.Value)
}

func TestShouldUpdateACustomer(t *testing.T) {
	Setup()

	err := customerService.Register(customer_dto.CustomerRegisterDto{
		Name: "Jhon Doe",
		CPF:  "45616278041",
	})
	assert.Nil(t, err)

	repo := customer_repository.NewCustomerRepository(test.DB.DB)
	customers, err := repo.FindAll()

	assert.Nil(t, err)
	assert.Equal(t, 1, len(customers))
	assert.NotNil(t, customers[0].Id.Value)
	assert.Equal(t, "Jhon Doe", customers[0].Name)
	assert.Equal(t, "45616278041", customers[0].CPF.Value)

	customerID := customers[0].Id.Value

	err = customerService.Update(customer_dto.CustomerUpdateDto{
		Id:   customerID,
		Name: "Jhon Doe 2",
	})
	assert.Nil(t, err)

	customer, err := repo.FindById(customerID)

	assert.Nil(t, err)
	assert.Equal(t, "Jhon Doe 2", customer.Name)
}
