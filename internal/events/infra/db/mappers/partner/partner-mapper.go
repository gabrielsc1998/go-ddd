package partner_mapper

import (
	partner_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/partner"
	partner_model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/partner"
)

type PartnerMapper struct {
}

func NewPartnerMapper() *PartnerMapper {
	return &PartnerMapper{}
}

func (mapper *PartnerMapper) ToModel(entity *partner_entity.Partner) *partner_model.Partner {
	return &partner_model.Partner{
		ID:   entity.Id.Value,
		Name: entity.Name,
	}
}

func (mapper *PartnerMapper) ToEntity(model *partner_model.Partner) *partner_entity.Partner {
	partner, _ := partner_entity.Create(partner_entity.PartnerCreateProps{
		Id:   model.ID,
		Name: model.Name,
	})
	return partner
}
