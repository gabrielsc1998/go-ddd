package application_service

import (
	domain_event_manager "github.com/gabrielsc1998/go-ddd/internal/common/domain/domain-event-manager"
	"github.com/gabrielsc1998/go-ddd/internal/common/domain/entity"
)

type ApplicationServiceInterface interface {
	Start()
	Fail()
	Run(aggregateRoots []*entity.AggregateRoot, callback func() error) error
}

type ApplicationService struct {
	domainEventManager domain_event_manager.DomainEventManagerInterface
}

func NewApplicationService(domainEventManager domain_event_manager.DomainEventManagerInterface) *ApplicationService {
	return &ApplicationService{
		domainEventManager: domainEventManager,
	}
}

func (a *ApplicationService) Start() {
}

func (a *ApplicationService) Fail() {
}

func (a *ApplicationService) Run(aggregateRoots []*entity.AggregateRoot, callback func() error) error {
	err := callback()
	if err != nil {
		a.Fail()
		return err
	}
	for _, aggregateRoot := range aggregateRoots {
		a.domainEventManager.PublishForDomainEvent(aggregateRoot)
		a.domainEventManager.PublishForIntegrationEvent(aggregateRoot)
	}
	return nil
}
