package partner_domain_events

import (
	"time"

	domain_event "github.com/gabrielsc1998/go-ddd/internal/common/domain/domain-event"
)

type PartnerCreatedDomainEvent struct {
	domain_event.DomainEvent
}

func NewPartnerCreatedDomainEvent(aggregateId string) *PartnerCreatedDomainEvent {
	return &PartnerCreatedDomainEvent{
		DomainEvent: domain_event.DomainEvent{
			Name:         "PartnerCreated",
			AggregateId:  aggregateId,
			OccuredOn:    time.Now(),
			EventVersion: 1,
		},
	}
}
