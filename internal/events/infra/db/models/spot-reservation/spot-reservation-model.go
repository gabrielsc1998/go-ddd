package spot_reservation_model

import (
	"time"

	spot_model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/spot"
	"gorm.io/gorm"
)

type SpotReservation struct {
	gorm.Model
	ID              string          `gorm:"primaryKey;type:varchar(36);"`
	CustomerId      string          `gorm:"type:varchar(36);"`
	SpotId          string          `gorm:"type:varchar(36);"`
	Spot            spot_model.Spot `gorm:"foreignKey:SpotId;references:ID"`
	ReservationDate time.Time       `gorm:"not null"`
}
