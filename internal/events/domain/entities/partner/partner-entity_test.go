package partner_entity

import (
	"testing"
	"time"

	partner_events "github.com/gabrielsc1998/go-ddd/internal/events/domain/events/partner"
	"github.com/stretchr/testify/assert"
)

func TestShouldReturnErrorWhenPassAnInvalidId(t *testing.T) {
	_, err := Create(PartnerCreateProps{
		Id:   "invalid",
		Name: "name",
	})
	assert.Error(t, err, "invalid id")
}

func TestShouldReturnErrorWhenPassAnEmptyName(t *testing.T) {
	_, err := Create(PartnerCreateProps{
		Id:   "fc50a094-edc0-400a-ac80-b728ebd0270d",
		Name: "",
	})
	assert.Error(t, err, "invalid name")
}

func TestShouldCreateAPartner(t *testing.T) {
	partner, err := Create(PartnerCreateProps{
		Id:   "",
		Name: "Name",
	})

	assert.NoError(t, err)
	assert.NotEmpty(t, partner.Id.Value)
	assert.Equal(t, "Name", partner.Name)

	partner, err = Create(PartnerCreateProps{
		Id:   "fc50a094-edc0-400a-ac80-b728ebd0270d",
		Name: "Name",
	})

	assert.NoError(t, err)
	assert.NotEmpty(t, partner.Id.Value)
	assert.Equal(t, "fc50a094-edc0-400a-ac80-b728ebd0270d", partner.Id.Value)
	assert.Equal(t, "Name", partner.Name)

	domainEvents := partner.AggregateRoot.GetEvents()
	assert.Len(t, domainEvents, 1)
	assert.Equal(t, "fc50a094-edc0-400a-ac80-b728ebd0270d", domainEvents[0].(*partner_events.PartnerCreatedEvent).AggregateId)
}

func TestShouldInitAnEvent(t *testing.T) {
	partner, err := Create(PartnerCreateProps{
		Id:   "",
		Name: "Name",
	})

	assert.NoError(t, err)

	date := time.Now()
	event, err := partner.InitEvent(PartnerInitEventCommand{
		Name:        "Event Name",
		Description: "Event Description",
		Date:        date,
	})

	assert.NoError(t, err)
	assert.NotEmpty(t, event.Id.Value)
	assert.Equal(t, "Event Name", event.Name)
	assert.Equal(t, "Event Description", event.Description)
}
