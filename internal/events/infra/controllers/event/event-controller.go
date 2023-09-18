package event_controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gabrielsc1998/go-ddd/internal/common/domain/value-objects/id"
	event_dto "github.com/gabrielsc1998/go-ddd/internal/events/application/dto/event"
	order_dto "github.com/gabrielsc1998/go-ddd/internal/events/application/dto/order"
	event_service "github.com/gabrielsc1998/go-ddd/internal/events/application/services/event"
	order_service "github.com/gabrielsc1998/go-ddd/internal/events/application/services/order"
	event_presenter "github.com/gabrielsc1998/go-ddd/internal/events/infra/presenter/event"
	order_presenter "github.com/gabrielsc1998/go-ddd/internal/events/infra/presenter/order"
	section_presenter "github.com/gabrielsc1998/go-ddd/internal/events/infra/presenter/section"
	spot_presenter "github.com/gabrielsc1998/go-ddd/internal/events/infra/presenter/spot"
	"github.com/go-chi/chi/v5"
)

type EventController struct {
	eventService event_service.EventService
	orderService order_service.OrderService
}

func NewEventController(eventService event_service.EventService, orderService order_service.OrderService) *EventController {
	return &EventController{
		eventService: eventService,
		orderService: orderService,
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
	eventIdParam := chi.URLParam(r, "event_id")
	eventId, err := id.NewID(eventIdParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	eventSections, err := e.eventService.FindSections(eventId.Value)
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
	eventIdParam := chi.URLParam(r, "event_id")
	eventId, err := id.NewID(eventIdParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var dto AddSectionInputDto
	err = json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = e.eventService.AddSection(event_dto.EventAddSectionDto{
		EventId:            eventId.Value,
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

func (e *EventController) PublishAll(w http.ResponseWriter, r *http.Request) {
	eventIdParam := chi.URLParam(r, "event_id")
	eventId, err := id.NewID(eventIdParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = e.eventService.PublishAll(eventId.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (e *EventController) UpdateSection(w http.ResponseWriter, r *http.Request) {
	eventIdParam := chi.URLParam(r, "event_id")
	eventId, err := id.NewID(eventIdParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sectionIdParam := chi.URLParam(r, "section_id")
	sectionId, err := id.NewID(sectionIdParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var dto UpdateSectionInputDto
	err = json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = e.eventService.UpdateSectionInformation(event_dto.EventUpdateSectionDto{
		EventId:     eventId.Value,
		SectionId:   sectionId.Value,
		Name:        dto.Name,
		Description: dto.Description,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (e *EventController) GetSectionSpots(w http.ResponseWriter, r *http.Request) {
	eventIdParam := chi.URLParam(r, "event_id")
	eventId, err := id.NewID(eventIdParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sectionIdParam := chi.URLParam(r, "section_id")
	sectionId, err := id.NewID(sectionIdParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sectionSpots, err := e.eventService.FindSpots(event_dto.EventFindSpotsDto{
		EventId:   eventId.Value,
		SectionId: sectionId.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sectionSpotsPresenter := make([]spot_presenter.SpotPresenter, len(sectionSpots))
	for i, spot := range sectionSpots {
		sectionSpotsPresenter[i] = spot_presenter.ToPresent(&spot)
	}
	json.NewEncoder(w).Encode(sectionSpotsPresenter)
}

func (e *EventController) UpdateLocation(w http.ResponseWriter, r *http.Request) {
	eventIdParam := chi.URLParam(r, "event_id")
	eventId, err := id.NewID(eventIdParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	sectionIdParam := chi.URLParam(r, "section_id")
	sectionId, err := id.NewID(sectionIdParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	spotIdParam := chi.URLParam(r, "spot_id")
	spotId, err := id.NewID(spotIdParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var dto UpdateLocationInputDto
	err = json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = e.eventService.UpdateLocation(event_dto.EventUpdateLocationDto{
		EventId:   eventId.Value,
		SectionId: sectionId.Value,
		SpotId:    spotId.Value,
		Location:  dto.Location,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (e *EventController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	eventIdParam := chi.URLParam(r, "event_id")
	eventId, err := id.NewID(eventIdParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var dto CreateOrderInputDto
	err = json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = e.orderService.Create(order_dto.OrderCreateDto{
		EventId:    eventId.Value,
		CustomerId: dto.CustomerId,
		SectionId:  dto.SectionId,
		SpotId:     dto.SpotId,
		CardToken:  dto.CardToken,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (e *EventController) ListOrders(w http.ResponseWriter, r *http.Request) {
	eventIdParam := chi.URLParam(r, "event_id")
	eventId, err := id.NewID(eventIdParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	order, err := e.orderService.List(eventId.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	orderPresenter := make([]order_presenter.OrderPresenter, len(order))
	for i, event := range order {
		orderPresenter[i] = order_presenter.ToPresent(event)
	}
	json.NewEncoder(w).Encode(orderPresenter)
}
