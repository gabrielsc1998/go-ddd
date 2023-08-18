package spot_entity

import (
	"github.com/gabrielsc1998/go-ddd/internal/common/domain/entity"
	"github.com/gabrielsc1998/go-ddd/internal/common/domain/value-objects/id"
)

type Spot struct {
	entity.Entity
	Location    string `json:"location"`
	IsReserved  bool   `json:"is_reserved"`
	IsPublished bool   `json:"is_published"`
}

type SpotCreateProps struct {
	Id          string `json:"id"`
	Location    string `json:"location"`
	IsReserved  bool   `json:"is_reserved"`
	IsPublished bool   `json:"is_published"`
}

func Create(props SpotCreateProps) (*Spot, error) {
	err := validate(props)
	if err != nil {
		return nil, err
	}
	eventId, _ := id.NewID(props.Id)
	return &Spot{
		Entity: entity.Entity{
			Id: eventId,
		},
		Location:    props.Location,
		IsReserved:  props.IsReserved,
		IsPublished: props.IsPublished,
	}, nil
}

func validate(props SpotCreateProps) error {
	_, err := id.NewID(props.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Spot) Publish() {
	s.IsPublished = true
}

func (s *Spot) Reserve() {
	s.IsReserved = true
}

func (s *Spot) Unreserve() {
	s.IsReserved = false
}

func (s *Spot) ChangeLocation(location string) {
	s.Location = location
}
