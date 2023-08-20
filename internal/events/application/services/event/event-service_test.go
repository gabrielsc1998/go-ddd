package event_service

import (
	"testing"
	"time"

	"github.com/gabrielsc1998/go-ddd/internal/common/tests"
	event_dto "github.com/gabrielsc1998/go-ddd/internal/events/application/dto/event"
	partner_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/partner"
	event_repository "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/repositories/event"
	partner_repository "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/repositories/partner"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var test *tests.Tests
var eventService EventService
var partnerRepository *partner_repository.PartnerRepository

func Setup() {
	test = tests.Setup()
	test.UOW.Register("EventRepository", func(db *gorm.DB) interface{} {
		repo := event_repository.NewEventRepository(db)
		return repo
	})
	partnerRepository = partner_repository.NewPartnerRepository(test.DB.DB)
	eventService = NewEventService(EventServiceProps{
		UOW:               test.UOW,
		EventRepository:   event_repository.NewEventRepository(test.DB.DB),
		PartnerRepository: partnerRepository,
	})
}

func TestShouldCreateAnEvent(t *testing.T) {
	Setup()

	partner, _ := partner_entity.Create(partner_entity.PartnerCreateProps{
		Id:   "",
		Name: "Jhon Doe",
	})
	err := partnerRepository.Add(partner)
	assert.Nil(t, err)

	date := time.Now()
	err = eventService.Create(event_dto.EventCreateDto{
		Name:        "Event 01",
		Description: "Event 01 description",
		Date:        date,
		PartnerId:   partner.Id.Value,
	})
	assert.Nil(t, err)

	repo := event_repository.NewEventRepository(test.DB.DB)
	event, err := repo.FindAll()

	assert.Nil(t, err)
	assert.Equal(t, 1, len(event))
	assert.NotNil(t, event[0].Id.Value)
	assert.Equal(t, "Event 01", event[0].Name)
	assert.Equal(t, "Event 01 description", event[0].Description)
	assert.Equal(t, date.Format("2006-01-02 15:04:05:000"), event[0].Date.Format("2006-01-02 15:04:05:000"))
	assert.Equal(t, partner.Id.Value, event[0].PartnerId.Value)
}

func TestShouldUpdateAnEvent(t *testing.T) {
	Setup()

	partner, _ := partner_entity.Create(partner_entity.PartnerCreateProps{
		Id:   "",
		Name: "Jhon Doe",
	})
	err := partnerRepository.Add(partner)
	assert.Nil(t, err)

	date := time.Now()
	err = eventService.Create(event_dto.EventCreateDto{
		Name:        "Event 01",
		Description: "Event 01 description",
		Date:        date,
		PartnerId:   partner.Id.Value,
	})
	assert.Nil(t, err)

	repo := event_repository.NewEventRepository(test.DB.DB)
	events, err := repo.FindAll()

	assert.Nil(t, err)
	assert.Equal(t, 1, len(events))
	assert.NotNil(t, events[0].Id.Value)
	assert.Equal(t, "Event 01", events[0].Name)
	assert.Equal(t, "Event 01 description", events[0].Description)
	assert.Equal(t, date.Format("2006-01-02 15:04:05:000"), events[0].Date.Format("2006-01-02 15:04:05:000"))
	assert.Equal(t, partner.Id.Value, events[0].PartnerId.Value)

	date2 := time.Now()
	err = eventService.Update(event_dto.EventUpdateDto{
		Id:          events[0].Id.Value,
		Name:        "Event 02",
		Description: "Event 02 description",
		Date:        date2,
	})
	assert.Nil(t, err)

	event, err := repo.FindById(events[0].Id.Value)

	assert.Nil(t, err)
	assert.Equal(t, "Event 02", event.Name)
	assert.Equal(t, "Event 02 description", event.Description)
	assert.Equal(t, date.Format("2006-01-02 15:04:05:000"), event.Date.Format("2006-01-02 15:04:05:000"))

	events, err = repo.FindAll()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(events))
}

func TestShouldGetAllEvents(t *testing.T) {
	Setup()

	partner, _ := partner_entity.Create(partner_entity.PartnerCreateProps{
		Id:   "",
		Name: "Jhon Doe",
	})
	err := partnerRepository.Add(partner)
	assert.Nil(t, err)

	date := time.Now()
	err = eventService.Create(event_dto.EventCreateDto{
		Name:        "Event 01",
		Description: "Event 01 description",
		Date:        date,
		PartnerId:   partner.Id.Value,
	})
	assert.Nil(t, err)

	err = eventService.Create(event_dto.EventCreateDto{
		Name:        "Event 02",
		Description: "Event 02 description",
		Date:        date,
		PartnerId:   partner.Id.Value,
	})
	assert.Nil(t, err)

	events, err := eventService.FindEvents()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(events))
	assert.Equal(t, "Event 01", events[0].Name)
	assert.Equal(t, partner.Id.Value, events[0].PartnerId.Value)
	assert.Equal(t, "Event 01 description", events[0].Description)
	assert.Equal(t, date.Format("2006-01-02 15:04:05:000"), events[0].Date.Format("2006-01-02 15:04:05:000"))
	assert.Equal(t, "Event 02", events[1].Name)
	assert.Equal(t, "Event 02 description", events[1].Description)
	assert.Equal(t, date.Format("2006-01-02 15:04:05:000"), events[1].Date.Format("2006-01-02 15:04:05:000"))
	assert.Equal(t, partner.Id.Value, events[1].PartnerId.Value)
}
