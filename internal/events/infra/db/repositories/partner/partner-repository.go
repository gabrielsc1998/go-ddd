package partner_repository

import (
	"gorm.io/gorm"

	entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/partner"
	mapper "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/mappers/partner"
	model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/partner"
)

type PartnerRepository struct {
	db     *gorm.DB
	mapper *mapper.PartnerMapper
}

func NewPartnerRepository(db *gorm.DB) *PartnerRepository {
	mapper := mapper.NewPartnerMapper()
	return &PartnerRepository{db: db, mapper: mapper}
}

func (r *PartnerRepository) Add(partner *entity.Partner) error {
	return r.db.Create(r.mapper.ToModel(partner)).Error
}

func (r *PartnerRepository) FindById(id string) (*entity.Partner, error) {
	var partner model.Partner
	err := r.db.Where("id = ?", id).First(&partner).Error
	return r.mapper.ToEntity(&partner), err
}

func (r *PartnerRepository) FindAll() ([]*entity.Partner, error) {
	var partners []*model.Partner
	err := r.db.Find(&partners).Error
	var partnersEntity []*entity.Partner
	for _, partner := range partners {
		partnersEntity = append(partnersEntity, r.mapper.ToEntity(partner))
	}
	return partnersEntity, err
}

func (r *PartnerRepository) Delete(partner *entity.Partner) error {
	return r.db.Delete(r.mapper.ToModel(partner)).Error
}
