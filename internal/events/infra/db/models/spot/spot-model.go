package spot_model

import (
	"gorm.io/gorm"
)

type Spot struct {
	gorm.Model
	ID          string `gorm:"primaryKey;type:varchar(36);"`
	Location    string `gorm:"not null"`
	IsReserved  bool   `gorm:"not null"`
	IsPublished bool   `gorm:"not null"`
	SectionId   string `gorm:"type:varchar(36);"`
}
