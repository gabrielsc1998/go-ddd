package id

import (
	"errors"

	"github.com/google/uuid"
)

type ID struct {
	Value string
}

func NewID(value string) (*ID, error) {
	if value == "" {
		return &ID{Value: uuid.New().String()}, nil
	}
	id := &ID{Value: value}
	err := id.Validate()
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (i *ID) Validate() error {
	_, err := uuid.Parse(i.Value)
	if err != nil {
		return errors.New("invalid id")
	}
	return nil
}
