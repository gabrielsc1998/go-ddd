package outbox

import (
	"time"

	"gorm.io/gorm"
)

type OutboxModel struct {
	gorm.Model
	Id               string     `gorm:"primaryKey;type:varchar(36);"`
	Data             []byte     `gorm:"type:blob;not null;"`
	State            int        `gorm:"not null;default:0;"`
	LockedBy         string     `gorm:"type:varchar(36);"`
	LockedOn         *time.Time `gorm:""`
	ProcessedOn      *time.Time `gorm:""`
	NumberOfAttempts int        `gorm:"not null;default:0;"`
	Error            string     `gorm:"type:varchar(1000);"`
}

func (OutboxModel) TableName() string {
	return "outbox"
}
