package partner_presenter

import partner_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/partner"

type PartnerPresenter struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func ToPresent(partner *partner_entity.Partner) PartnerPresenter {
	return PartnerPresenter{
		Id:   partner.Id.Value,
		Name: partner.Name,
	}
}
