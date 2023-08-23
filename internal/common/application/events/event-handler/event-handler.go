package event_handler

import (
	"sync"

	domain_event "github.com/gabrielsc1998/go-ddd/internal/common/domain/domain-event"
)

type EventHandler struct {
	callback func(event domain_event.DomainEvent, wg *sync.WaitGroup)
}

func NewEventHandler(callback func(event domain_event.DomainEvent, wg *sync.WaitGroup)) *EventHandler {
	return &EventHandler{
		callback: callback,
	}
}

func (eh *EventHandler) Handle(event domain_event.DomainEvent, wg *sync.WaitGroup) {
	eh.callback(event, wg)
}
