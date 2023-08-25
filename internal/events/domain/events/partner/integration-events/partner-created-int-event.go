package partner_int_events

import (
	"time"

	domain_event "github.com/gabrielsc1998/go-ddd/internal/common/domain/domain-event"
	integration_event "github.com/gabrielsc1998/go-ddd/internal/common/domain/integration-event"
)

type PartnerCreatedIntegrationEvent struct {
	integration_event.IntegrationEvent
}

func GetName() string {
	return "PartnerCreatedInt"
}

func NewPartnerCreatedEvent(domainEvent *domain_event.DomainEvent) *PartnerCreatedIntegrationEvent {
	payload := struct {
		AggregateId string
		Name        string
	}{
		AggregateId: domainEvent.AggregateId,
		Name:        domainEvent.Name,
	}
	return &PartnerCreatedIntegrationEvent{
		IntegrationEvent: integration_event.IntegrationEvent{
			Name:         GetName(),
			Payload:      payload,
			OccuredOn:    time.Now(),
			EventVersion: 1,
		},
	}
}
