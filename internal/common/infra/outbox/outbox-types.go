package outbox

import "gorm.io/gorm"

type DtoAddInOutbox struct {
	Payload []byte
}

type OutboxInterface interface {
	Add(input DtoAddInOutbox, lockedBy string) error
	Lock(id string) error
	Unlock(id string) error
	MarkAsProcessed(id string) error
	MarkAsFailed(id string) error
	GetUnprocessedOutbox() ([]OutboxModel, error)
}

type Outbox struct {
	db *gorm.DB
}

var OutboxStatus = struct {
	UNPROCESSED int
	PROCESSED   int
	LOCKED      int
	FAILED      int
}{
	UNPROCESSED: 0,
	PROCESSED:   1,
	LOCKED:      2,
	FAILED:      3,
}
