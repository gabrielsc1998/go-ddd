package spot_reservation_entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestShouldCreateASpotReservation(t *testing.T) {
	date := time.Now()
	spotReservation, err := Create(SpotReservationCreateProps{
		Id:              "",
		SpotId:          "c3957510-ca0f-4e2a-ac30-010912165618",
		CustomerId:      "e71d6dd3-8113-4bc5-b8f6-ace1d5034175",
		ReservationDate: date,
	})

	assert.NoError(t, err)
	assert.NotEmpty(t, spotReservation.Id.Value)
	assert.Equal(t, spotReservation.SpotId.Value, "c3957510-ca0f-4e2a-ac30-010912165618")
	assert.Equal(t, spotReservation.CustomerId.Value, "e71d6dd3-8113-4bc5-b8f6-ace1d5034175")
	assert.Equal(t, spotReservation.ReservationDate, date)
}

func TestShouldChangeReservation(t *testing.T) {
	date := time.Now()
	spotReservation, err := Create(SpotReservationCreateProps{
		Id:              "",
		SpotId:          "c3957510-ca0f-4e2a-ac30-010912165618",
		CustomerId:      "e71d6dd3-8113-4bc5-b8f6-ace1d5034175",
		ReservationDate: date,
	})

	assert.NoError(t, err)
	assert.Equal(t, spotReservation.CustomerId.Value, "e71d6dd3-8113-4bc5-b8f6-ace1d5034175")
	assert.Equal(t, spotReservation.ReservationDate, date)

	date2 := date.AddDate(0, 0, 1)
	spotReservation.ChangeReservation(SpotReservationCommandChangeReservation{
		CustomerId:      "a4c08f45-86cd-477b-b85d-7447cc0f632e",
		ReservationDate: date2,
	})

	assert.Equal(t, spotReservation.CustomerId.Value, "a4c08f45-86cd-477b-b85d-7447cc0f632e")
	assert.Equal(t, spotReservation.ReservationDate, date2)
}
