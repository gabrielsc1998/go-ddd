package event_controller

import (
	"encoding/json"
	"net/http"
	"time"

	event_dto "github.com/gabrielsc1998/go-ddd/internal/events/application/dto/event"
	event_service "github.com/gabrielsc1998/go-ddd/internal/events/application/services/event"
	event_presenter "github.com/gabrielsc1998/go-ddd/internal/events/infra/presenter/event"
	section_presenter "github.com/gabrielsc1998/go-ddd/internal/events/infra/presenter/section"
	"github.com/go-chi/chi/v5"
)

type EventController struct {
	eventService event_service.EventService
}

func NewEventController(eventService event_service.EventService) *EventController {
	return &EventController{
		eventService: eventService,
	}
}

func (e *EventController) CreateEvent(w http.ResponseWriter, r *http.Request) {
	var dto CreateEventInputDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	date, _ := time.Parse("2006-01-02T15:04:05.000Z", dto.Date)
	err = e.eventService.Create(event_dto.EventCreateDto{
		Name:        dto.Name,
		Description: dto.Description,
		Date:        date,
		PartnerId:   dto.PartnerId,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (e *EventController) ListEvents(w http.ResponseWriter, r *http.Request) {
	events, err := e.eventService.FindEvents()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	eventsPresenter := make([]event_presenter.EventPresenter, len(events))
	for i, event := range events {
		eventsPresenter[i] = event_presenter.ToPresent(event)
	}
	json.NewEncoder(w).Encode(eventsPresenter)
}

func (e *EventController) FindEventSections(w http.ResponseWriter, r *http.Request) {
	eventId := chi.URLParam(r, "event_id")
	eventSections, err := e.eventService.FindSections(eventId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	eventSectionsPresenter := make([]section_presenter.SectionPresenter, len(eventSections))
	for i, section := range eventSections {
		eventSectionsPresenter[i] = section_presenter.ToPresent(&section)
	}
	json.NewEncoder(w).Encode(eventSectionsPresenter)
}

func (e *EventController) AddSection(w http.ResponseWriter, r *http.Request) {
	eventId := chi.URLParam(r, "event_id")
	var dto AddSectionInputDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = e.eventService.AddSection(event_dto.EventAddSectionDto{
		EventId:            eventId,
		Name:               dto.Name,
		Description:        dto.Description,
		Date:               time.Now(),
		IsPublished:        false,
		TotalSpots:         dto.TotalSpots,
		TotalSpotsReserved: 0,
		Price:              dto.Price,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
