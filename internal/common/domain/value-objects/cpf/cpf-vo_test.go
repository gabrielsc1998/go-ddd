package cpf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnAnErrorWhenPassInvalidCPF(t *testing.T) {
	arrange := []string{"00000000000", "12345678999", "1111", "", " "}
	for _, value := range arrange {
		_, err := NewCPF(value)
		assert.Error(t, err, "invalid CPF")
	}
}

func TestShouldNotReturnErrorWhenPassAnValidCPF(t *testing.T) {
	arrange := []string{"263.464.030-79", "26346403079"}
	for _, value := range arrange {
		cpf, err := NewCPF(value)
		assert.NoError(t, err)
		assert.NotEmpty(t, cpf)
		assert.NoError(t, cpf.Validate())
		assert.Equal(t, cpf.Clear(value), cpf.Value)
		assert.Equal(t, cpf.Format(value), cpf.Formatted)
	}
}
