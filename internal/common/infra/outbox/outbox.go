package outbox

import (
	"time"

	cronjob "github.com/gabrielsc1998/go-ddd/internal/common/application/cron-job"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func NewOutbox(db *gorm.DB, handle func(*[]OutboxModel) error) *Outbox {
	db.AutoMigrate(&OutboxModel{})
	outbox := &Outbox{
		db: db,
	}
	cronJob := cronjob.NewCronJob(5, func() {
		outboxes, _ := outbox.GetUnprocessedOutbox()
		handle(outboxes)
	})
	cronJob.Start()
	return outbox
}

func (o *Outbox) Add(input DtoAddInOutbox) error {
	return o.db.Create(&OutboxModel{
		Id:               uuid.New().String(),
		Data:             input.Payload,
		State:            OutboxStatus.UNPROCESSED,
		NumberOfAttempts: 0,
	}).Error
}

func (o *Outbox) getOutboxData(id string) (OutboxModel, error) {
	outboxData := OutboxModel{}
	err := o.db.Where("id = ?", id).First(&outboxData).Error
	return outboxData, err
}

func (o *Outbox) Lock(id string, lockedBy string) error {
	outboxData, err := o.getOutboxData(id)
	if err != nil {
		return err
	}
	lockedOn := time.Now()
	outboxData.LockedOn = &lockedOn
	outboxData.LockedBy = lockedBy
	outboxData.State = OutboxStatus.LOCKED
	return o.db.Save(&outboxData).Error
}

func (o *Outbox) Unlock(id string) error {
	outboxData, err := o.getOutboxData(id)
	if err != nil {
		return err
	}
	outboxData.LockedBy = ""
	outboxData.State = OutboxStatus.PROCESSED
	return o.db.Save(&outboxData).Error
}

func (o *Outbox) MarkAsProcessed(id string) error {
	outboxData, err := o.getOutboxData(id)
	if err != nil {
		return err
	}
	processedOn := time.Now()
	outboxData.ProcessedOn = &processedOn
	outboxData.LockedBy = ""
	outboxData.State = OutboxStatus.PROCESSED
	return o.db.Save(&outboxData).Error
}

func (o *Outbox) MarkAsFailed(id string) error {
	outboxData, err := o.getOutboxData(id)
	if err != nil {
		return err
	}
	outboxData.LockedBy = ""
	outboxData.State = OutboxStatus.FAILED
	return o.db.Save(&outboxData).Error
}

func (o *Outbox) GetUnprocessedOutbox() (*[]OutboxModel, error) {
	outboxesData := &[]OutboxModel{}
	err := o.db.Where("state = ?", OutboxStatus.UNPROCESSED).Find(&outboxesData).Error
	if err != nil {
		return nil, err
	}
	return outboxesData, nil
}
