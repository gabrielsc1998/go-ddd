package id

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnAnErrorWhenPassInvalidId(t *testing.T) {
	_, err := NewID("invalid id")
	assert.EqualError(t, err, "invalid id")
}

func TestShouldNotReturnErrorWhenPassAnValidId(t *testing.T) {
	id, err := NewID("55b58442-0dae-4eaa-a871-dde00af8c1b4")
	assert.NoError(t, err)
	assert.NotEmpty(t, id)
	assert.NoError(t, id.Validate())
	assert.Equal(t, "55b58442-0dae-4eaa-a871-dde00af8c1b4", id.Value)
}

func TestShouldReturnAnIdWhenNotPassValue(t *testing.T) {
	id, err := NewID("")
	assert.NoError(t, err)
	assert.NotEmpty(t, id)
	assert.NoError(t, id.Validate())
}
