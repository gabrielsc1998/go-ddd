package spot_reservation_mapper

import (
	spot_reservation_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/spot-reservation"
	spot_reservation_model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/spot-reservation"
)

type SpotReservationMapper struct {
}

func NewSpotReservationMapper() *SpotReservationMapper {
	return &SpotReservationMapper{}
}

func (mapper *SpotReservationMapper) ToModel(entity *spot_reservation_entity.SpotReservation) *spot_reservation_model.SpotReservation {
	return &spot_reservation_model.SpotReservation{
		ID:              entity.Id.Value,
		CustomerId:      entity.CustomerId.Value,
		SpotId:          entity.SpotId.Value,
		ReservationDate: entity.ReservationDate,
	}
}

func (mapper *SpotReservationMapper) ToEntity(model *spot_reservation_model.SpotReservation) *spot_reservation_entity.SpotReservation {
	section, _ := spot_reservation_entity.Create(spot_reservation_entity.SpotReservationCreateProps{
		Id:              model.ID,
		CustomerId:      model.CustomerId,
		SpotId:          model.SpotId,
		ReservationDate: model.ReservationDate,
	})
	return section
}
