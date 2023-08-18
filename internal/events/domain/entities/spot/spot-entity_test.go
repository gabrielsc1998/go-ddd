package spot_entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCreateASpot(t *testing.T) {
	spot, err := Create(SpotCreateProps{
		Id:          "",
		Location:    "Location",
		IsReserved:  false,
		IsPublished: false,
	})

	assert.NoError(t, err)
	assert.NotEmpty(t, spot.Id.Value)
	assert.Equal(t, "Location", spot.Location)
	assert.Equal(t, false, spot.IsReserved)
	assert.Equal(t, false, spot.IsPublished)
}

func TestShouldPublish(t *testing.T) {
	spot, _ := Create(SpotCreateProps{
		Id:          "",
		Location:    "Location",
		IsReserved:  false,
		IsPublished: false,
	})
	spot.Publish()
	assert.True(t, spot.IsPublished)
}
