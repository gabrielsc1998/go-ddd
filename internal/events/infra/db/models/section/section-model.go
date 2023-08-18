package section_model

import (
	"time"

	spot_model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/spot"
	"gorm.io/gorm"
)

type Section struct {
	gorm.Model
	ID                 string            `gorm:"primary_key;type:varchar(36);"`
	Name               string            `gorm:"not null"`
	Description        string            `gorm:"not null"`
	Date               time.Time         `gorm:"not null"`
	IsPublished        bool              `gorm:"not null"`
	TotalSpots         int               `gorm:"not null"`
	TotalSpotsReserved int               `gorm:"not null"`
	Price              float64           `gorm:"not null"`
	EventId            string            `gorm:"type:varchar(36);"`
	Spots              []spot_model.Spot `gorm:"foreignKey:SectionId"`
}
