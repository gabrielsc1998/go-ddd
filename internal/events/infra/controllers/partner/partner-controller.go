package partner_controller

import (
	"encoding/json"
	"net/http"

	partner_dto "github.com/gabrielsc1998/go-ddd/internal/events/application/dto/partner"
	partner_service "github.com/gabrielsc1998/go-ddd/internal/events/application/services/partner"
	partner_presenter "github.com/gabrielsc1998/go-ddd/internal/events/infra/presenter/partner"
)

type PartnerController struct {
	eventService partner_service.PartnerService
}

func NewPartnerController(eventService partner_service.PartnerService) *PartnerController {
	return &PartnerController{
		eventService: eventService,
	}
}

func (p *PartnerController) CreatePartner(w http.ResponseWriter, r *http.Request) {
	var dto CreatePartnerInputDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = p.eventService.Register(partner_dto.PartnerRegisterDto{
		Name: dto.Name,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (p *PartnerController) ListPartners(w http.ResponseWriter, r *http.Request) {
	partners, err := p.eventService.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	partnersPresenter := make([]partner_presenter.PartnerPresenter, len(partners))
	for i, partner := range partners {
		partnersPresenter[i] = partner_presenter.ToPresent(partner)
	}
	json.NewEncoder(w).Encode(partnersPresenter)
}
