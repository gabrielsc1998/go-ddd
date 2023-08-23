package domain_event_manager

import (
	event_dispatcher "github.com/gabrielsc1998/go-ddd/internal/common/application/events/event-dispatcher"
	"github.com/gabrielsc1998/go-ddd/internal/common/domain/entity"
)

type DomainEventManagerInterface interface {
	Register(event string, handler event_dispatcher.EventHandlerInterface)
	Publish(aggregateRoot *entity.AggregateRoot)
}

type DomainEventManager struct {
	eventDispatcher event_dispatcher.EventDispatcherInterface
}

func NewDomainEventManager(eventDispatcher event_dispatcher.EventDispatcherInterface) *DomainEventManager {
	return &DomainEventManager{
		eventDispatcher: eventDispatcher,
	}
}

func (d *DomainEventManager) Register(event string, handler event_dispatcher.EventHandlerInterface) {
	d.eventDispatcher.Register(event, handler)
}

func (d *DomainEventManager) Publish(aggregateRoot *entity.AggregateRoot) {
	for _, event := range aggregateRoot.GetEvents() {
		d.eventDispatcher.Dispatch(event)
	}
}
