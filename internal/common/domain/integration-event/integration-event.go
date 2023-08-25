package integration_event

import "time"

type IntegrationEvent struct {
	Name         string
	Payload      interface{}
	OccuredOn    time.Time
	EventVersion int
}

func (i *IntegrationEvent) GetName() string {
	return i.Name
}
