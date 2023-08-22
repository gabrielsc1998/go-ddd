package event_entity

import (
	"errors"
	"time"

	"github.com/gabrielsc1998/go-ddd/internal/common/domain/entity"
	"github.com/gabrielsc1998/go-ddd/internal/common/domain/value-objects/id"
	section_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/section"
)

type Event struct {
	entity.AggregateRoot
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	Date               time.Time `json:"date"`
	IsPublished        bool      `json:"is_published"`
	TotalSpots         int       `json:"total_spots"`
	TotalSpotsReserved int       `json:"total_spots_reserved"`
	PartnerId          *id.ID    `json:"partner_id"`
	Sections           []section_entity.Section
}

type EventCreateProps struct {
	Id                 string    `json:"id"`
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	Date               time.Time `json:"date"`
	IsPublished        bool      `json:"is_published"`
	TotalSpots         int       `json:"total_spots"`
	TotalSpotsReserved int       `json:"total_spots_reserved"`
	PartnerId          string    `json:"partner_id"`
}

type EventCommandChangeSectionInfo struct {
	SectionId   string `json:"section_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type EventCommandChangeLocation struct {
	SectionId string `json:"section_id"`
	SpotId    string `json:"spot_id"`
	Location  string `json:"location"`
}

type EventCommandReserveSpot struct {
	SectionId string `json:"section_id"`
	SpotId    string `json:"spot_id"`
}

func Create(props EventCreateProps) (*Event, error) {
	err := validate(props)
	if err != nil {
		return nil, err
	}
	eventId, _ := id.NewID(props.Id)
	partnerId, _ := id.NewID(props.PartnerId)
	return &Event{
		AggregateRoot: entity.AggregateRoot{
			Id: eventId,
		},
		Name:               props.Name,
		Description:        props.Description,
		Date:               props.Date,
		IsPublished:        props.IsPublished,
		TotalSpots:         props.TotalSpots,
		TotalSpotsReserved: props.TotalSpotsReserved,
		PartnerId:          partnerId,
	}, nil
}

func validate(props EventCreateProps) error {
	_, err := id.NewID(props.Id)
	if err != nil {
		return err
	}
	_, err = id.NewID(props.PartnerId)
	if err != nil {
		return err
	}
	if props.Name == "" {
		return errors.New("invalid name")
	}
	return nil
}

func (e *Event) AddSection(props section_entity.SectionCreateProps) (*section_entity.Section, error) {
	section, err := section_entity.Create(section_entity.SectionCreateProps{
		Id:                 props.Id,
		Name:               props.Name,
		Description:        props.Description,
		Date:               props.Date,
		IsPublished:        props.IsPublished,
		TotalSpots:         props.TotalSpots,
		TotalSpotsReserved: props.TotalSpotsReserved,
		Price:              props.Price,
		Spots:              props.Spots,
	})
	if err != nil {
		return nil, err
	}
	e.Sections = append(e.Sections, *section)
	e.TotalSpots += section.TotalSpots
	return section, nil
}

func (e *Event) UpdateSectionInformation(command EventCommandChangeSectionInfo) error {
	section := &section_entity.Section{}
	for i := range e.Sections {
		if e.Sections[i].Id.Value == command.SectionId {
			section = &e.Sections[i]
			break
		}
	}
	if command.Name != "" && command.Name != section.Name {
		section.ChangeName(command.Name)
	}
	if command.Description != "" && command.Description != section.Description {
		section.ChangeDescription(command.Description)
	}
	return nil
}

func (e *Event) Publish() {
	e.IsPublished = true
}

func (e *Event) PublishAll() {
	e.Publish()
	for i := range e.Sections {
		e.Sections[i].PublishAll()
	}
}

func (e *Event) ChangeName(name string) {
	e.Name = name
}

func (e *Event) ChangeDescription(description string) {
	e.Description = description
}

func (e *Event) ChangeDate(date time.Time) {
	e.Date = date
}

func (e *Event) ChangeLocation(command EventCommandChangeLocation) error {
	section, err := e.GetSection(command.SectionId)
	if err != nil {
		return err
	}
	section.ChangeLocation(section_entity.SectionCommandChangeLocation{
		SpotId:   command.SpotId,
		Location: command.Location,
	})
	return nil
}

func (e *Event) AllowReserveSpot(command EventCommandReserveSpot) bool {
	if e.IsPublished == false {
		return false
	}
	section, err := e.GetSection(command.SectionId)
	if err != nil {
		return false
	}
	return section.AllowReserveSpot(command.SpotId)
}

func (e *Event) ReserveSpot(command EventCommandReserveSpot) error {
	section, err := e.GetSection(command.SectionId)
	if err != nil {
		return err
	}
	err = section.ReserveSpot(command.SpotId)
	if err != nil {
		return err
	}
	e.TotalSpotsReserved++
	return nil
}

func (e *Event) GetSection(sectionId string) (*section_entity.Section, error) {
	for i := range e.Sections {
		if e.Sections[i].Id.Value == sectionId {
			return &e.Sections[i], nil

		}
	}
	return nil, errors.New("section not found")
}
