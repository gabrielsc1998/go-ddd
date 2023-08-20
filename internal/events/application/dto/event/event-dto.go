package event_dto

import "time"

type EventCreateDto struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	PartnerId   string    `json:"partner_id"`
}

type EventUpdateDto struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

type EventAddSectionDto struct {
	EventId            string    `json:"event_id"`
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	Date               time.Time `json:"date"`
	IsPublished        bool      `json:"is_published"`
	TotalSpots         int       `json:"total_spots"`
	TotalSpotsReserved int       `json:"total_spots_reserved"`
	Price              float64   `json:"price"`
}

type EventUpdateSectionDto struct {
	EventId     string `json:"event_id"`
	SectionId   string `json:"section_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type EventFindSpotsDto struct {
	EventId   string `json:"event_id"`
	SectionId string `json:"section_id"`
}

type EventUpdateLocationDto struct {
	EventId   string `json:"event_id"`
	SectionId string `json:"section_id"`
	SpotId    string `json:"spot_id"`
	Location  string `json:"location"`
}
