package event_domain_events

import (
	"time"

	domain_event "github.com/gabrielsc1998/go-ddd/internal/common/domain/domain-event"
)

type EventCreatedDomainEvent struct {
	domain_event.DomainEvent
}

func NewEventCreatedDomainEvent(aggregateId string) *EventCreatedDomainEvent {
	return &EventCreatedDomainEvent{
		DomainEvent: domain_event.DomainEvent{
			Name:         "EventCreated",
			AggregateId:  aggregateId,
			OccuredOn:    time.Now(),
			EventVersion: 1,
		},
	}
}
