package valueobjects

import "errors"

type ValueObjectInterface interface {
	Validate() error
}

type ValueObject struct {
	Value string
}

func (v *ValueObject) Validate() error {
	return errors.New("Method not implemented")
}
