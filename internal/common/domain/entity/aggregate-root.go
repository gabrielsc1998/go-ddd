package entity

import (
	domain_event "github.com/gabrielsc1998/go-ddd/internal/common/domain/domain-event"
	"github.com/gabrielsc1998/go-ddd/internal/common/domain/value-objects/id"
)

type AggregateRootInterface interface {
	AddEvent(event *domain_event.DomainEvent)
	ClearEvents()
	GetEvents() []*domain_event.DomainEvent
}

type AggregateRoot struct {
	Id     *id.ID
	events []*domain_event.DomainEvent
}

func (a *AggregateRoot) AddEvent(event *domain_event.DomainEvent) {
	a.events = append(a.events, event)
}

func (a *AggregateRoot) ClearEvents() {
	a.events = []*domain_event.DomainEvent{}
}

func (a *AggregateRoot) GetEvents() []*domain_event.DomainEvent {
	return a.events
}
