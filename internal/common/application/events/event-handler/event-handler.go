package event_handler

import (
	"sync"
)

type EventHandler struct {
	callback func(event interface{}, wg *sync.WaitGroup)
}

func NewEventHandler(callback func(event interface{}, wg *sync.WaitGroup)) *EventHandler {
	return &EventHandler{
		callback: callback,
	}
}

func (eh *EventHandler) Handle(event interface{}, wg *sync.WaitGroup) {
	eh.callback(event, wg)
}
