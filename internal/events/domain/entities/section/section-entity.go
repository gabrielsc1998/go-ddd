package section_entity

import (
	"errors"
	"time"

	"github.com/gabrielsc1998/go-ddd/internal/common/domain/entity"
	"github.com/gabrielsc1998/go-ddd/internal/common/domain/value-objects/id"
	spot_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/spot"
)

type Section struct {
	entity.Entity
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	Date               time.Time `json:"date"`
	IsPublished        bool      `json:"is_published"`
	TotalSpots         int       `json:"total_spots"`
	TotalSpotsReserved int       `json:"total_spots_reserved"`
	Price              float64   `json:"price"`
	Spots              []spot_entity.Spot
}

type SectionCreateProps struct {
	Id                 string    `json:"id"`
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	Date               time.Time `json:"date"`
	IsPublished        bool      `json:"is_published"`
	TotalSpots         int       `json:"total_spots"`
	TotalSpotsReserved int       `json:"total_spots_reserved"`
	Price              float64   `json:"price"`
	Spots              []spot_entity.Spot
}

type SectionCommandChangeLocation struct {
	SpotId   string `json:"spot_id"`
	Location string `json:"location"`
}

func Create(props SectionCreateProps) (*Section, error) {
	err := validate(props)
	if err != nil {
		return nil, err
	}
	eventId, _ := id.NewID(props.Id)
	section := &Section{
		Entity: entity.Entity{
			Id: eventId,
		},
		Name:               props.Name,
		Description:        props.Description,
		Date:               props.Date,
		IsPublished:        props.IsPublished,
		TotalSpots:         props.TotalSpots,
		TotalSpotsReserved: props.TotalSpotsReserved,
		Price:              props.Price,
	}
	if len(props.Spots) > 0 {
		section.Spots = props.Spots
	} else {
		section.createSpots()
	}
	return section, nil
}

func (s *Section) createSpots() error {
	for i := 0; i < s.TotalSpots; i++ {
		_, err := s.AddSpot(spot_entity.SpotCreateProps{
			Id:          "",
			Location:    "",
			IsReserved:  false,
			IsPublished: false,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func validate(props SectionCreateProps) error {
	_, err := id.NewID(props.Id)
	if err != nil {
		return err
	}
	if props.Name == "" {
		return errors.New("invalid name")
	}
	return nil
}

func (s *Section) AddSpot(props spot_entity.SpotCreateProps) (*spot_entity.Spot, error) {
	spot, err := spot_entity.Create(props)
	if err != nil {
		return nil, err
	}
	s.Spots = append(s.Spots, *spot)
	return spot, nil
}

func (s *Section) Publish() {
	s.IsPublished = true
}

func (s *Section) PublishAll() {
	s.Publish()
	for i := 0; i < s.TotalSpots; i++ {
		s.Spots[i].Publish()
	}
}

func (s *Section) ChangeName(newName string) error {
	err := validate(SectionCreateProps{
		Id:                 s.Id.Value,
		Name:               newName,
		Description:        s.Description,
		Date:               s.Date,
		IsPublished:        s.IsPublished,
		TotalSpots:         s.TotalSpots,
		TotalSpotsReserved: s.TotalSpotsReserved,
		Price:              s.Price,
		Spots:              s.Spots,
	})
	if err != nil {
		return err
	}
	s.Name = newName
	return nil
}

func (s *Section) ChangeDescription(newDescription string) error {
	err := validate(SectionCreateProps{
		Id:                 s.Id.Value,
		Name:               s.Name,
		Description:        newDescription,
		Date:               s.Date,
		IsPublished:        s.IsPublished,
		TotalSpots:         s.TotalSpots,
		TotalSpotsReserved: s.TotalSpotsReserved,
		Price:              s.Price,
		Spots:              s.Spots,
	})
	if err != nil {
		return err
	}
	s.Description = newDescription
	return nil
}

func (s *Section) ChangeLocation(command SectionCommandChangeLocation) error {
	spot, err := s.getSpot(command.SpotId)
	if err != nil {
		return err
	}
	spot.ChangeLocation(command.Location)
	return nil
}

func (s *Section) AllowReserveSpot(spotId string) bool {
	if !s.IsPublished {
		return false
	}
	spot, err := s.getSpot(spotId)
	if err != nil {
		return false
	}
	if spot.IsReserved || !spot.IsPublished {
		return false
	}
	return true
}

func (s *Section) ReserveSpot(spotId string) error {
	spot, err := s.getSpot(spotId)
	if err != nil {
		return err
	}
	spot.Reserve()
	s.TotalSpotsReserved++
	return nil
}

func (s *Section) getSpot(spotId string) (*spot_entity.Spot, error) {
	for i := range s.Spots {
		if s.Spots[i].Id.Value == spotId {
			return &s.Spots[i], nil
		}
	}
	return nil, errors.New("spot not found")
}
