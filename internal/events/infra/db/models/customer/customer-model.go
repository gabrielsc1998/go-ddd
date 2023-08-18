package customer_model

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	ID   string `gorm:"primary_key;type:varchar(36);"`
	Name string `gorm:"not null"`
	CPF  string `gorm:"not null"`
}
