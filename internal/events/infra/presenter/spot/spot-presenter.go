package spot_presenter

import spot_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/spot"

type SpotPresenter struct {
	Id          string `json:"id"`
	Location    string `json:"location"`
	IsReserved  bool   `json:"is_reserved"`
	IsPublished bool   `json:"is_published"`
}

func ToPresent(spot *spot_entity.Spot) SpotPresenter {
	return SpotPresenter{
		Id:          spot.Id.Value,
		Location:    spot.Location,
		IsReserved:  spot.IsReserved,
		IsPublished: spot.IsPublished,
	}
}
