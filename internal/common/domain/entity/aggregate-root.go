package entity

import "github.com/gabrielsc1998/go-ddd/internal/common/domain/value-objects/id"

type AggregateRootInterface interface {
}

type AggregateRoot struct {
	Id     *id.ID
	events []interface{}
}

func (a *AggregateRoot) AddEvent(event interface{}) {
	a.events = append(a.events, event)
}

func (a *AggregateRoot) ClearEvents() {
	a.events = []interface{}{}
}

func (a *AggregateRoot) GetEvents() []interface{} {
	return a.events
}
