package event_service

import (
	"context"

	unit_of_work "github.com/gabrielsc1998/go-ddd/internal/common/infra/db/unit-of-work"
	event_dto "github.com/gabrielsc1998/go-ddd/internal/events/application/dto/event"
	event_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/event"
	partner_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/partner"
	section_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/section"
	spot_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/spot"
	event_repository "github.com/gabrielsc1998/go-ddd/internal/events/domain/repositories/event"
	partner_repository "github.com/gabrielsc1998/go-ddd/internal/events/domain/repositories/partner"
)

type EventService struct {
	uow               *unit_of_work.Uow
	eventRepository   event_repository.EventRepositoryInterface
	partnerRepository partner_repository.PartnerRepositoryInterface
}

type EventServiceProps struct {
	UOW               *unit_of_work.Uow
	EventRepository   event_repository.EventRepositoryInterface
	PartnerRepository partner_repository.PartnerRepositoryInterface
}

func NewEventService(props EventServiceProps) EventService {
	return EventService{
		uow:               props.UOW,
		eventRepository:   props.EventRepository,
		partnerRepository: props.PartnerRepository,
	}
}

func (s *EventService) getEventRepository() (event_repository.EventRepositoryInterface, error) {
	ctx := context.Background()
	repo, err := s.uow.GetRepository(ctx, "EventRepository")
	if err != nil {
		return nil, err
	}
	eventRepository := repo.(event_repository.EventRepositoryInterface)
	return eventRepository, nil
}

func (s *EventService) Create(input event_dto.EventCreateDto) error {
	partner, err := s.partnerRepository.FindById(input.PartnerId)
	if err != nil {
		return err
	}
	event, err := partner.InitEvent(partner_entity.PartnerInitEventCommand{
		Name:        input.Name,
		Description: input.Description,
		Date:        input.Date,
	})
	if err != nil {
		return err
	}
	eventRepository, err := s.getEventRepository()
	err = s.uow.Do(s.uow.GetCtx(), func(uow *unit_of_work.Uow) error {
		err = eventRepository.Add(event)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (s *EventService) Update(input event_dto.EventUpdateDto) error {
	eventRepository, err := s.getEventRepository()
	if err != nil {
		return err
	}

	event, err := eventRepository.FindById(input.Id)
	if err != nil {
		return err
	}
	if input.Name != "" {
		event.ChangeName(input.Name)
	}
	if input.Description != "" {
		event.ChangeDescription(input.Description)
	}
	if !input.Date.IsZero() {
		event.ChangeDate(input.Date)
	}
	err = s.uow.Do(s.uow.GetCtx(), func(uow *unit_of_work.Uow) error {
		err = eventRepository.Add(event)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (s *EventService) AddSection(input event_dto.EventAddSectionDto) error {
	eventRepository, err := s.getEventRepository()
	if err != nil {
		return err
	}
	event, err := eventRepository.FindById(input.EventId)
	if err != nil {
		return err
	}
	event.AddSection(section_entity.SectionCreateProps{
		Id:                 "",
		Name:               input.Name,
		Description:        input.Description,
		Date:               input.Date,
		IsPublished:        input.IsPublished,
		TotalSpots:         input.TotalSpots,
		TotalSpotsReserved: input.TotalSpotsReserved,
		Price:              input.Price,
	})
	err = s.uow.Do(s.uow.GetCtx(), func(uow *unit_of_work.Uow) error {
		err = eventRepository.Add(event)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (s *EventService) UpdateSectionInformation(input event_dto.EventUpdateSectionDto) error {
	event, err := s.eventRepository.FindById(input.EventId)
	if err != nil {
		return err
	}
	event.UpdateSectionInformation(event_entity.EventCommandChangeSectionInfo{
		SectionId:   input.SectionId,
		Name:        input.Name,
		Description: input.Description,
	})
	eventRepository, err := s.getEventRepository()
	err = s.uow.Do(s.uow.GetCtx(), func(uow *unit_of_work.Uow) error {
		err = eventRepository.Add(event)
		if err != nil {
			return err
		}
		return nil
	})
	return nil
}

func (s *EventService) FindEvents() ([]*event_entity.Event, error) {
	events, err := s.eventRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (s *EventService) FindSections(eventId string) ([]section_entity.Section, error) {
	event, err := s.eventRepository.FindById(eventId)
	if err != nil {
		return nil, err
	}
	return event.Sections, nil
}

func (s *EventService) FindSpots(input event_dto.EventFindSpotsDto) ([]spot_entity.Spot, error) {
	event, err := s.eventRepository.FindById(input.EventId)
	if err != nil {
		return nil, err
	}
	for _, section := range event.Sections {
		if section.Id.Value == input.SectionId {
			return section.Spots, nil
		}
	}
	return nil, nil
}

func (s *EventService) UpdateLocation(input event_dto.EventUpdateLocationDto) error {
	event, err := s.eventRepository.FindById(input.EventId)
	if err != nil {
		return err
	}
	err = event.ChangeLocation(event_entity.EventCommandChangeLocation{
		SectionId: input.SectionId,
		SpotId:    input.SpotId,
		Location:  input.Location,
	})
	if err != nil {
		return err
	}
	eventRepository, err := s.getEventRepository()
	err = s.uow.Do(s.uow.GetCtx(), func(uow *unit_of_work.Uow) error {
		err = eventRepository.Add(event)
		if err != nil {
			return err
		}
		return nil
	})
	return nil
}

func (s *EventService) PublishAll(eventId string) error {
	event, err := s.eventRepository.FindById(eventId)
	if err != nil {
		return err
	}
	event.PublishAll()
	eventRepository, err := s.getEventRepository()
	err = s.uow.Do(s.uow.GetCtx(), func(uow *unit_of_work.Uow) error {
		err = eventRepository.Add(event)
		if err != nil {
			return err
		}
		return nil
	})
	return nil
}
