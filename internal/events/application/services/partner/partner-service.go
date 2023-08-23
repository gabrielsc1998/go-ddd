package partner_service

import (
	"context"

	application_service "github.com/gabrielsc1998/go-ddd/internal/common/application/application-service"
	"github.com/gabrielsc1998/go-ddd/internal/common/domain/entity"
	unit_of_work "github.com/gabrielsc1998/go-ddd/internal/common/infra/db/unit-of-work"
	partner_dto "github.com/gabrielsc1998/go-ddd/internal/events/application/dto/partner"
	partner_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/partner"
	partner_repository "github.com/gabrielsc1998/go-ddd/internal/events/domain/repositories/partner"
)

type PartnerService struct {
	uow                *unit_of_work.Uow
	partnerRepository  partner_repository.PartnerRepositoryInterface
	applicationService application_service.ApplicationServiceInterface
}

type PartnerServiceProps struct {
	UOW                *unit_of_work.Uow
	PartnerRepository  partner_repository.PartnerRepositoryInterface
	ApplicationService application_service.ApplicationServiceInterface
}

func NewPartnerService(props PartnerServiceProps) PartnerService {
	return PartnerService{
		uow:                props.UOW,
		partnerRepository:  props.PartnerRepository,
		applicationService: props.ApplicationService,
	}
}

func (p *PartnerService) getPartnerRepository() (partner_repository.PartnerRepositoryInterface, error) {
	ctx := context.Background()
	repo, err := p.uow.GetRepository(ctx, "PartnerRepository")
	if err != nil {
		return nil, err
	}
	partnerRepository := repo.(partner_repository.PartnerRepositoryInterface)
	return partnerRepository, nil
}

func (p *PartnerService) Register(input partner_dto.PartnerRegisterDto) error {
	partner, _ := partner_entity.Create(partner_entity.PartnerCreateProps{
		Id:   "",
		Name: input.Name,
	})
	aggregateRoots := make([]*entity.AggregateRoot, 0)
	aggregateRoots = append(aggregateRoots, &partner.AggregateRoot)
	return p.applicationService.Run(aggregateRoots, func() error {
		partnerRepository, err := p.getPartnerRepository()
		err = p.uow.Do(p.uow.GetCtx(), func(uow *unit_of_work.Uow) error {
			err = partnerRepository.Add(partner)
			if err != nil {
				return err
			}
			return nil
		})
		return nil
	})
}

func (p *PartnerService) Update(input partner_dto.PartnerUpdateDto) error {
	partnerRepository, err := p.getPartnerRepository()
	if err != nil {
		return err
	}
	partner, err := partnerRepository.FindById(input.Id)
	if err != nil {
		return err
	}
	if input.Name != "" {
		partner.ChangeName(input.Name)
	}
	err = p.uow.Do(p.uow.GetCtx(), func(uow *unit_of_work.Uow) error {
		err = partnerRepository.Add(partner)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (p *PartnerService) List() ([]*partner_entity.Partner, error) {
	partnerRepository, err := p.getPartnerRepository()
	if err != nil {
		return nil, err
	}
	partners, err := partnerRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return partners, nil
}
