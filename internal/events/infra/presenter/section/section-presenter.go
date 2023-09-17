package section_presenter

import section_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/section"

type SectionPresenter struct {
	Id                 string  `json:"id"`
	Name               string  `json:"name"`
	Description        string  `json:"description"`
	Date               string  `json:"date"`
	IsPublished        bool    `json:"is_published"`
	TotalSpots         int     `json:"total_spots"`
	TotalSpotsReserved int     `json:"total_spots_reserved"`
	Price              float64 `json:"price"`
}

func ToPresent(section *section_entity.Section) SectionPresenter {
	return SectionPresenter{
		Id:                 section.Id.Value,
		Name:               section.Name,
		Description:        section.Description,
		Date:               section.Date.Format("2006-01-02 15:04:05.000"),
		IsPublished:        section.IsPublished,
		TotalSpots:         section.TotalSpots,
		TotalSpotsReserved: section.TotalSpotsReserved,
		Price:              section.Price,
	}
}
