package event_model

import (
	"time"

	partner_model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/partner"
	section_model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/section"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	ID                 string                  `gorm:"primary_key;type:varchar(36);"`
	Name               string                  `gorm:"not null"`
	Description        string                  `gorm:"not null"`
	Date               time.Time               `gorm:"not null"`
	IsPublished        bool                    `gorm:"not null"`
	TotalSpots         int                     `gorm:"not null"`
	TotalSpotsReserved int                     `gorm:"not null"`
	PartnerId          string                  `gorm:"type:varchar(36);"`
	Partner            partner_model.Partner   `gorm:"foreignKey:PartnerId"`
	Sections           []section_model.Section `gorm:"foreignKey:EventId"`
}
