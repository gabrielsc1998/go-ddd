package domain_event

import "time"

type DomainEvent struct {
	AggregateId  string
	OccuredOn    time.Time
	EventVersion int
}
