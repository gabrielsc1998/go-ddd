package spot_reservation_model

import (
	"time"

	"gorm.io/gorm"
)

type SpotReservation struct {
	gorm.Model
	ID              string    `gorm:"primaryKey;type:varchar(36);"`
	CustomerId      string    `gorm:"type:varchar(36);"`
	SpotId          string    `gorm:"type:varchar(36);"`
	ReservationDate time.Time `gorm:"not null"`
}
