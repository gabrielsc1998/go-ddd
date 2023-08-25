package domain_event

import "time"

type DomainEvent struct {
	Name         string
	AggregateId  string
	OccuredOn    time.Time
	EventVersion int
}

func (d *DomainEvent) GetName() string {
	return d.Name
}
