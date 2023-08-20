package partner_repository

import entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/partner"

type PartnerRepositoryInterface interface {
	Add(partner *entity.Partner) error
	FindById(id string) (*entity.Partner, error)
	FindAll() ([]*entity.Partner, error)
	Delete(partner *entity.Partner) error
}
