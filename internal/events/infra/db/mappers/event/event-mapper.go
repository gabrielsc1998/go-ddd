package event_mapper

import (
	event_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/event"
	section_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/section"
	section_mapper "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/mappers/section"
	event_model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/event"
	section_model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/section"
)

type EventMapper struct {
}

func NewEventMapper() *EventMapper {
	return &EventMapper{}
}

func (mapper *EventMapper) ToModel(entity *event_entity.Event) *event_model.Event {
	sections := make([]section_model.Section, 0)
	sectionMapper := section_mapper.NewSectionMapper()
	for _, sectionEntity := range entity.Sections {
		section := sectionMapper.ToModel(&sectionEntity)
		sections = append(sections, *section)
	}
	return &event_model.Event{
		ID:                 entity.Id.Value,
		Name:               entity.Name,
		Description:        entity.Description,
		Date:               entity.Date,
		IsPublished:        entity.IsPublished,
		TotalSpots:         entity.TotalSpots,
		TotalSpotsReserved: entity.TotalSpotsReserved,
		PartnerId:          entity.PartnerId.Value,
		Sections:           sections,
	}
}

func (mapper *EventMapper) ToEntity(model *event_model.Event) *event_entity.Event {
	event, _ := event_entity.Create(event_entity.EventCreateProps{
		Id:                 model.ID,
		Name:               model.Name,
		Description:        model.Description,
		Date:               model.Date,
		IsPublished:        model.IsPublished,
		TotalSpots:         model.TotalSpots,
		TotalSpotsReserved: model.TotalSpotsReserved,
		PartnerId:          model.PartnerId,
	})
	for _, sectionModel := range model.Sections {
		sectionMapper := section_mapper.NewSectionMapper()
		section := sectionMapper.ToEntity(&sectionModel)
		event.AddSection(section_entity.SectionCreateProps{
			Id:                 section.Id.Value,
			Name:               section.Name,
			Description:        section.Description,
			Date:               section.Date,
			IsPublished:        section.IsPublished,
			TotalSpots:         section.TotalSpots,
			TotalSpotsReserved: section.TotalSpotsReserved,
			Price:              section.Price,
			Spots:              section.Spots,
		})
	}
	return event
}
