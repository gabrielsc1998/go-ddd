package event_presenter

import event_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/event"

type EventPresenter struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        string `json:"date"`
	PartnerId   string `json:"partner_id"`
}

func ToPresent(event *event_entity.Event) EventPresenter {
	return EventPresenter{
		Id:          event.Id.Value,
		Name:        event.Name,
		Description: event.Description,
		Date:        event.Date.Format("2006-01-02 15:04:05.000"),
		PartnerId:   event.PartnerId.Value,
	}
}
