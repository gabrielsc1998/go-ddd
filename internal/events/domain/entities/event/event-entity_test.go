package event_entity

import (
	"testing"
	"time"

	section_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/section"
	"github.com/stretchr/testify/assert"
)

func TestShouldCreateAEvent(t *testing.T) {
	date := time.Now()
	event, err := Create(EventCreateProps{
		Id:                 "",
		Name:               "Name",
		Description:        "Description",
		Date:               date,
		IsPublished:        false,
		TotalSpots:         0,
		TotalSpotsReserved: 0,
	})

	event.AddSection(section_entity.SectionCreateProps{
		Id:                 "",
		Name:               "Name",
		Description:        "Description",
		Date:               date,
		IsPublished:        false,
		TotalSpots:         10,
		TotalSpotsReserved: 0,
		Price:              10.0,
	})

	assert.NoError(t, err)
	assert.NotEmpty(t, event.Id.Value)
	assert.Equal(t, "Name", event.Name)
	assert.Equal(t, "Description", event.Description)
	assert.Equal(t, date, event.Date)
	assert.Equal(t, false, event.IsPublished)
	assert.Equal(t, 10, event.TotalSpots)
	assert.Equal(t, 0, event.TotalSpotsReserved)
	assert.Equal(t, 1, len(event.Sections))
	assert.Equal(t, 10, len(event.Sections[0].Spots))
}

func TestShouldPublishAll(t *testing.T) {
	date := time.Now()
	event, err := Create(EventCreateProps{
		Id:                 "",
		Name:               "Name",
		Description:        "Description",
		Date:               date,
		IsPublished:        false,
		TotalSpots:         1,
		TotalSpotsReserved: 0,
	})

	event.AddSection(section_entity.SectionCreateProps{
		Id:                 "",
		Name:               "Name",
		Description:        "Description",
		Date:               date,
		IsPublished:        false,
		TotalSpots:         2,
		TotalSpotsReserved: 0,
		Price:              10.0,
	})

	event.AddSection(section_entity.SectionCreateProps{
		Id:                 "",
		Name:               "Name 2",
		Description:        "Description 2",
		Date:               date,
		IsPublished:        false,
		TotalSpots:         2,
		TotalSpotsReserved: 0,
		Price:              10.0,
	})

	assert.NoError(t, err)

	assert.Equal(t, false, event.IsPublished)
	assert.Equal(t, false, event.Sections[0].IsPublished)
	assert.Equal(t, false, event.Sections[0].Spots[0].IsPublished)

	event.PublishAll()

	assert.Equal(t, true, event.IsPublished)
	assert.Equal(t, true, event.Sections[0].IsPublished)
	assert.Equal(t, true, event.Sections[0].Spots[0].IsPublished)
	assert.Equal(t, true, event.Sections[0].Spots[1].IsPublished)
	assert.Equal(t, true, event.Sections[1].IsPublished)
	assert.Equal(t, true, event.Sections[1].Spots[0].IsPublished)
	assert.Equal(t, true, event.Sections[1].Spots[1].IsPublished)
}
