package event_repository

import entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/event"

type EventRepositoryInterface interface {
	Add(event *entity.Event) error
	FindById(id string) (*entity.Event, error)
	FindAll() ([]*entity.Event, error)
	Delete(event *entity.Event) error
}
