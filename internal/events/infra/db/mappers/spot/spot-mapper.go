package spot_mapper

import (
	spot_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/spot"
	spot_model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/spot"
)

type SpotMapper struct {
}

func NewSpotMapper() *SpotMapper {
	return &SpotMapper{}
}

func (mapper *SpotMapper) ToModel(entity *spot_entity.Spot) *spot_model.Spot {
	return &spot_model.Spot{
		ID:          entity.Id.Value,
		Location:    entity.Location,
		IsReserved:  entity.IsReserved,
		IsPublished: entity.IsPublished,
	}
}

func (mapper *SpotMapper) ToEntity(model *spot_model.Spot) *spot_entity.Spot {
	section, _ := spot_entity.Create(spot_entity.SpotCreateProps{
		Id:          model.ID,
		Location:    model.Location,
		IsReserved:  model.IsReserved,
		IsPublished: model.IsPublished,
	})
	return section
}
