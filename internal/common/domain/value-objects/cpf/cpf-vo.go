package cpf

import (
	"errors"
	"regexp"

	vo "github.com/gabrielsc1998/go-ddd/internal/common/domain/value-objects"
	"github.com/klassmann/cpfcnpj"
)

type CPF struct {
	vo.ValueObject
	Formatted string
}

func NewCPF(value string) (*CPF, error) {
	receivedValue := cpfcnpj.Clean(value)
	cpf := &CPF{ValueObject: vo.ValueObject{Value: receivedValue}, Formatted: ""}
	err := cpf.Validate()
	cpf.Formatted = cpf.Format(receivedValue)
	if err != nil {
		return nil, err
	}
	return cpf, nil
}

func (c *CPF) Validate() error {
	if c.Value == "00000000000" {
		return errors.New("invalid CPF")
	}
	result := cpfcnpj.ValidateCPF(c.Value)
	if !result {
		return errors.New("invalid CPF")
	}
	return nil
}

func (c *CPF) Clear(value string) string {
	return cpfcnpj.Clean(value)
}

func (c *CPF) Format(value string) string {
	e, _ := regexp.Compile(cpfcnpj.CPFFormatPattern)
	return e.ReplaceAllString(value, "$1.$2.$3-$4")
}
