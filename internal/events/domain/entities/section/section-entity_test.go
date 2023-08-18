package section_entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestShouldCreateASection(t *testing.T) {
	date := time.Now()
	section, err := Create(SectionCreateProps{
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
	assert.NotEmpty(t, section.Id.Value)
	assert.Equal(t, "Name", section.Name)
	assert.Equal(t, "Description", section.Description)
	assert.Equal(t, date, section.Date)
	assert.Equal(t, false, section.IsPublished)
	assert.Equal(t, 10, section.TotalSpots)
	assert.Equal(t, 0, section.TotalSpotsReserved)
	assert.Equal(t, 10.0, section.Price)
	assert.Equal(t, 10, len(section.Spots))
}

func TestShouldPublishAll(t *testing.T) {
	date := time.Now()
	section, err := Create(SectionCreateProps{
		Id:                 "",
		Name:               "Name",
		Description:        "Description",
		Date:               date,
		IsPublished:        false,
		TotalSpots:         2,
		TotalSpotsReserved: 0,
		Price:              10.0,
	})

	section.PublishAll()

	assert.NoError(t, err)
	assert.Equal(t, true, section.IsPublished)
	assert.Equal(t, true, section.Spots[0].IsPublished)
	assert.Equal(t, true, section.Spots[1].IsPublished)
}
