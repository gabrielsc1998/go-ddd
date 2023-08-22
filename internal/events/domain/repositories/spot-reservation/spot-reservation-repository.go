package spot_reservation_repository

import entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/spot-reservation"

type SpotReservationRepositoryInterface interface {
	Add(spotReservation *entity.SpotReservation) error
	FindById(id string) (*entity.SpotReservation, error)
	FindAll() ([]*entity.SpotReservation, error)
	Delete(spotReservation *entity.SpotReservation) error
}
