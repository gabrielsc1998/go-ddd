package spot_reservation_repository

import (
	"gorm.io/gorm"

	entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/spot-reservation"
	mapper "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/mappers/spot-reservation"
	model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/spot-reservation"
)

type SpotReservationRepository struct {
	db     *gorm.DB
	mapper *mapper.SpotReservationMapper
}

func NewSpotReservationRepository(db *gorm.DB) *SpotReservationRepository {
	mapper := mapper.NewSpotReservationMapper()
	return &SpotReservationRepository{db: db, mapper: mapper}
}

func (r *SpotReservationRepository) Add(spotReservation *entity.SpotReservation) error {
	spotReservationExists, _ := r.FindById(spotReservation.Id.Value)
	if spotReservationExists != nil {
		return r.db.Updates(r.mapper.ToModel(spotReservation)).Error
	}
	return r.db.Create(r.mapper.ToModel(spotReservation)).Error
}

func (r *SpotReservationRepository) FindById(id string) (*entity.SpotReservation, error) {
	var spotReservation model.SpotReservation
	err := r.db.Where("id = ?", id).First(&spotReservation).Error
	if err != nil {
		return nil, err
	}
	return r.mapper.ToEntity(&spotReservation), nil
}

func (r *SpotReservationRepository) FindAll() ([]*entity.SpotReservation, error) {
	var spotReservations []*model.SpotReservation
	err := r.db.Find(&spotReservations).Error
	var spotReservationsEntity []*entity.SpotReservation
	for _, spotReservation := range spotReservations {
		spotReservationsEntity = append(spotReservationsEntity, r.mapper.ToEntity(spotReservation))
	}
	return spotReservationsEntity, err
}

func (r *SpotReservationRepository) Delete(spotReservation *entity.SpotReservation) error {
	return r.db.Delete(r.mapper.ToModel(spotReservation)).Error
}
