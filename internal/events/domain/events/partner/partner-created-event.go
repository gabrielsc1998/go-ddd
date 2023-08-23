package partner_events

import (
	"time"

	domain_event "github.com/gabrielsc1998/go-ddd/internal/common/domain/domain-event"
)

type PartnerCreatedEvent struct {
	domain_event.DomainEvent
}

func NewPartnerCreatedEvent(aggregateId string) *PartnerCreatedEvent {
	return &PartnerCreatedEvent{
		DomainEvent: domain_event.DomainEvent{
			AggregateId:  aggregateId,
			OccuredOn:    time.Now(),
			EventVersion: 1,
		},
	}
}
