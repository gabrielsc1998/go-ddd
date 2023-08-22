package partner_service

import (
	"context"

	unit_of_work "github.com/gabrielsc1998/go-ddd/internal/common/infra/db/unit-of-work"
	partner_dto "github.com/gabrielsc1998/go-ddd/internal/events/application/dto/partner"
	partner_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/partner"
	partner_repository "github.com/gabrielsc1998/go-ddd/internal/events/domain/repositories/partner"
)

type PartnerService struct {
	uow               *unit_of_work.Uow
	partnerRepository partner_repository.PartnerRepositoryInterface
}

type PartnerServiceProps struct {
	UOW               *unit_of_work.Uow
	PartnerRepository partner_repository.PartnerRepositoryInterface
}

func NewPartnerService(props PartnerServiceProps) PartnerService {
	return PartnerService{
		uow:               props.UOW,
		partnerRepository: props.PartnerRepository,
	}
}

func (s *PartnerService) getPartnerRepository() (partner_repository.PartnerRepositoryInterface, error) {
	ctx := context.Background()
	repo, err := s.uow.GetRepository(ctx, "PartnerRepository")
	if err != nil {
		return nil, err
	}
	partnerRepository := repo.(partner_repository.PartnerRepositoryInterface)
	return partnerRepository, nil
}

func (s *PartnerService) Register(input partner_dto.PartnerRegisterDto) error {
	partner, _ := partner_entity.Create(partner_entity.PartnerCreateProps{
		Id:   "",
		Name: input.Name,
	})
	partnerRepository, err := s.getPartnerRepository()
	err = s.uow.Do(s.uow.GetCtx(), func(uow *unit_of_work.Uow) error {
		err = partnerRepository.Add(partner)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (s *PartnerService) Update(input partner_dto.PartnerUpdateDto) error {
	partnerRepository, err := s.getPartnerRepository()
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
	err = s.uow.Do(s.uow.GetCtx(), func(uow *unit_of_work.Uow) error {
		err = partnerRepository.Add(partner)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
