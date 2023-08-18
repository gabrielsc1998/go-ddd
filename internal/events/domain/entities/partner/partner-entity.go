package partner_entity

import (
	"errors"
	"time"

	"github.com/gabrielsc1998/go-ddd/internal/common/domain/entity"
	"github.com/gabrielsc1998/go-ddd/internal/common/domain/value-objects/id"
	event_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/event"
)

type Partner struct {
	entity.Entity
	Name string `json:"name"`
}

type PartnerCreateProps struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type PartnerInitEventCommand struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

func Create(props PartnerCreateProps) (*Partner, error) {
	err := validate(props)
	if err != nil {
		return nil, err
	}
	customerId, _ := id.NewID(props.Id)
	return &Partner{
		Entity: entity.Entity{
			Id: customerId,
		},
		Name: props.Name,
	}, nil
}

func validate(props PartnerCreateProps) error {
	_, err := id.NewID(props.Id)
	if err != nil {
		return err
	}
	if props.Name == "" {
		return errors.New("invalid name")
	}
	return nil
}

func (c *Partner) ChangeName(newName string) error {
	err := validate(PartnerCreateProps{
		Id:   c.Id.Value,
		Name: newName,
	})
	if err != nil {
		return err
	}
	c.Name = newName
	return nil
}

func (c *Partner) InitEvent(command PartnerInitEventCommand) (*event_entity.Event, error) {
	event, err := event_entity.Create(event_entity.EventCreateProps{
		Id:          "",
		Name:        command.Name,
		Description: command.Description,
		Date:        command.Date,
		PartnerId:   c.Id.Value,
	})
	if err != nil {
		return nil, err
	}
	return event, nil
}
