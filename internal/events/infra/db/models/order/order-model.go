package order_model

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID          string  `gorm:"primaryKey;type:varchar(36);"`
	CustomerId  string  `gorm:"type:varchar(36);"`
	Amount      float64 `gorm:"not null"`
	EventSpotId string  `gorm:"type:varchar(36);"`
	Status      int     `gorm:"not null;default:0"`
}
