package customer_entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnErrorWhenPassAnInvalidId(t *testing.T) {
	_, err := Create(CustomerCreateProps{
		Id:   "invalid",
		Name: "name",
		CPF:  "263.464.030-79",
	})
	assert.Error(t, err, "invalid id")
}

func TestShouldReturnErrorWhenPassAnEmptyName(t *testing.T) {
	_, err := Create(CustomerCreateProps{
		Id:   "fc50a094-edc0-400a-ac80-b728ebd0270d",
		Name: "",
		CPF:  "263.464.030-79",
	})
	assert.Error(t, err, "invalid name")
}

func TestShouldReturnErrorWhenPassAnInvalidCPF(t *testing.T) {
	_, err := Create(CustomerCreateProps{
		Id:   "fc50a094-edc0-400a-ac80-b728ebd0270d",
		Name: "Name",
		CPF:  "000000",
	})
	assert.Error(t, err, "invalid CPF")
}

func TestShouldCreateACustomer(t *testing.T) {
	customer, err := Create(CustomerCreateProps{
		Id:   "",
		Name: "Name",
		CPF:  "263.464.030-79",
	})

	assert.NoError(t, err)
	assert.NotEmpty(t, customer.Id.Value)
	assert.Equal(t, "Name", customer.Name)
	assert.Equal(t, "26346403079", customer.CPF.Value)
	assert.Equal(t, "263.464.030-79", customer.CPF.Formatted)

	customer, err = Create(CustomerCreateProps{
		Id:   "fc50a094-edc0-400a-ac80-b728ebd0270d",
		Name: "Name",
		CPF:  "263.464.030-79",
	})

	assert.NoError(t, err)
	assert.NotEmpty(t, customer.Id.Value)
	assert.Equal(t, "fc50a094-edc0-400a-ac80-b728ebd0270d", customer.Id.Value)
	assert.Equal(t, "Name", customer.Name)
	assert.Equal(t, "26346403079", customer.CPF.Value)
	assert.Equal(t, "263.464.030-79", customer.CPF.Formatted)
}
