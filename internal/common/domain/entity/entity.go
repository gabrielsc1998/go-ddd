package entity

import (
	"github.com/gabrielsc1998/go-ddd/internal/common/domain/value-objects/id"
)

type EntityInterface interface {
	// ToJSON() ([]byte, error)
}

type Entity struct {
	Id *id.ID `json:"id"`
}

// func (e *Entity) ToJSON() error {
// 	return errors.New("Method not implemented")
// }
