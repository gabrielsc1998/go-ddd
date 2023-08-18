package entity

import "github.com/gabrielsc1998/go-ddd/internal/common/domain/value-objects/id"

type AggregateRootInterface interface {
}

type AggregateRoot struct {
	Id *id.ID `json:"id"`
}
