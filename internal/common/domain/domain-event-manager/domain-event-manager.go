package domain_event_manager

import (
	event_dispatcher "github.com/gabrielsc1998/go-ddd/internal/common/application/events/event-dispatcher"
	"github.com/gabrielsc1998/go-ddd/internal/common/domain/entity"
)

type DomainEventManagerInterface interface {
	RegisterForDomainEvent(event string, handler event_dispatcher.EventHandlerInterface)
	PublishForDomainEvent(aggregateRoot *entity.AggregateRoot)
	RegisterForIntegrationEvent(event string, handler event_dispatcher.EventHandlerInterface)
	PublishForIntegrationEvent(aggregateRoot *entity.AggregateRoot)
}

type DomainEventManager struct {
	eventDomainDispatcher      event_dispatcher.EventDispatcherInterface
	eventIntegrationDispatcher event_dispatcher.EventDispatcherInterface
}

func NewDomainEventManager() *DomainEventManager {
	return &DomainEventManager{
		eventDomainDispatcher:      event_dispatcher.NewEventDispatcher(),
		eventIntegrationDispatcher: event_dispatcher.NewEventDispatcher(),
	}
}

func (d *DomainEventManager) RegisterForDomainEvent(event string, handler event_dispatcher.EventHandlerInterface) {
	d.eventDomainDispatcher.Register(event, handler)
}

func (d *DomainEventManager) PublishForDomainEvent(aggregateRoot *entity.AggregateRoot) {
	for _, event := range aggregateRoot.GetEvents() {
		d.eventDomainDispatcher.Dispatch(event)
	}
}

func (d *DomainEventManager) RegisterForIntegrationEvent(event string, handler event_dispatcher.EventHandlerInterface) {
	d.eventIntegrationDispatcher.Register(event, handler)
}

func (d *DomainEventManager) PublishForIntegrationEvent(aggregateRoot *entity.AggregateRoot) {
	for _, event := range aggregateRoot.GetEvents() {
		d.eventIntegrationDispatcher.Dispatch(event)
	}
}
