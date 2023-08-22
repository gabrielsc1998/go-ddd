package section_mapper

import (
	section_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/section"
	spot_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/spot"
	spot_mapper "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/mappers/spot"
	section_model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/section"
	spot_model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/spot"
)

type SectionMapper struct {
}

func NewSectionMapper() *SectionMapper {
	return &SectionMapper{}
}

func (mapper *SectionMapper) ToModel(entity *section_entity.Section) *section_model.Section {
	spots := make([]spot_model.Spot, 0)
	spotMapper := spot_mapper.NewSpotMapper()
	for _, spotEntity := range entity.Spots {
		spot := spotMapper.ToModel(&spotEntity)
		spots = append(spots, *spot)
	}
	return &section_model.Section{
		ID:                 entity.Id.Value,
		Name:               entity.Name,
		Description:        entity.Description,
		Date:               entity.Date,
		IsPublished:        entity.IsPublished,
		TotalSpots:         entity.TotalSpots,
		TotalSpotsReserved: entity.TotalSpotsReserved,
		Spots:              spots,
		Price:              entity.Price,
	}
}

func (mapper *SectionMapper) ToEntity(model *section_model.Section) *section_entity.Section {
	spots := make([]spot_entity.Spot, 0)
	spotMapper := spot_mapper.NewSpotMapper()
	for _, spotModel := range model.Spots {
		spot := spotMapper.ToEntity(&spotModel)
		spots = append(spots, *spot)
	}
	section, _ := section_entity.Create(section_entity.SectionCreateProps{
		Id:                 model.ID,
		Name:               model.Name,
		Description:        model.Description,
		Date:               model.Date,
		IsPublished:        model.IsPublished,
		TotalSpots:         model.TotalSpots,
		TotalSpotsReserved: model.TotalSpotsReserved,
		Spots:              spots,
		Price:              model.Price,
	})

	return section
}
