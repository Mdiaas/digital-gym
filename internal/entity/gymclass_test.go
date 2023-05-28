package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGymClass(t *testing.T) {
	gymClass, err := NewGymClass("Yoga", "https://www.youtube.com/watch?v=OuCvUDVoX3M")
	assert.Nil(t, err)
	assert.NotNil(t, gymClass)
	assert.Equal(t, "Yoga", gymClass.Name)
	assert.Equal(t, "https://www.youtube.com/watch?v=OuCvUDVoX3M", gymClass.Link)
	assert.NotEmpty(t, gymClass.ID)
}

func TestGymClassWithoutName(t *testing.T) {
	gymClass, err := NewGymClass("", "https://www.youtube.com/watch?v=OuCvUDVoX3M")
	assert.Nil(t, gymClass)
	assert.Equal(t, ErrNameIsRequired, err)
}
func TestGymClassWithoutLink(t *testing.T) {
	gymClass, err := NewGymClass("Yoga", "")
	assert.Nil(t, gymClass)
	assert.Equal(t, ErrLinkIsRequired, err)
}
func TestGymClassValidate(t *testing.T) {
	gymClass, err := NewGymClass("Yoga", "https://www.youtube.com/watch?v=OuCvUDVoX3M")
	assert.Nil(t, err)
	assert.NotNil(t, gymClass)
	assert.Nil(t, gymClass.Validate())
}
