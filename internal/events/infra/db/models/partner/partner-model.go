package partner_model

import (
	"gorm.io/gorm"
)

type Partner struct {
	gorm.Model
	ID   string `gorm:"primary_key;type:varchar(36);"`
	Name string `gorm:"not null"`
}
