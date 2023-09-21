package order_model

import (
	event_model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/event"
	spot_model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/spot"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID          string            `gorm:"primaryKey;type:varchar(36);"`
	CustomerId  string            `gorm:"type:varchar(36);"`
	Amount      float64           `gorm:"not null"`
	EventId     string            `gorm:"type:varchar(36);"`
	Event       event_model.Event `gorm:"foreignKey:EventId"`
	EventSpotId string            `gorm:"type:varchar(36);"`
	Spot        spot_model.Spot   `gorm:"foreignKey:EventSpotId"`
	Status      int               `gorm:"not null;default:0"`
}
