package event_repository

import (
	"gorm.io/gorm"

	entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/event"
	mapper "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/mappers/event"
	model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/event"
)

type EventRepository struct {
	db     *gorm.DB
	mapper *mapper.EventMapper
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	mapper := mapper.NewEventMapper()
	return &EventRepository{db: db, mapper: mapper}
}

func (r *EventRepository) Add(event *entity.Event) error {
	eventExists, _ := r.FindById(event.Id.Value)
	if eventExists != nil {
		return r.update(event)
	}
	return r.db.Create(r.mapper.ToModel(event)).Error
}

func (r *EventRepository) update(event *entity.Event) error {
	return r.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(r.mapper.ToModel(event)).Error
}

func (r *EventRepository) FindById(id string) (*entity.Event, error) {
	var event model.Event
	err := r.db.Where("id = ?", id).Preload("Sections").Preload("Sections.Spots").First(&event).Error
	return r.mapper.ToEntity(&event), err
}

func (r *EventRepository) FindAll() ([]*entity.Event, error) {
	var events []model.Event
	err := r.db.Find(&events).Order("created_at ASC").Error
	var eventsEntity []*entity.Event
	for _, event := range events {
		eventsEntity = append(eventsEntity, r.mapper.ToEntity(&event))
	}
	return eventsEntity, err
}

func (r *EventRepository) Delete(event *entity.Event) error {
	return r.db.Delete(r.mapper.ToModel(event)).Error
}
