package entity

import (
	"testing"

	"github.com/mdiaas/goapi/pkg/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Mateus Dias", "maugusto.diaas@gmail.com", "123456", "37386390041", true)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "Mateus Dias", user.FullName)
	assert.NotEmpty(t, user.Password)
	assert.NotEqual(t, "123456", user.Password)
	assert.Equal(t, "maugusto.diaas@gmail.com", user.Email)
	assert.Equal(t, entity.CPF("37386390041"), user.CPF)
	assert.Equal(t, true, user.IsAdmin)
}

func TestUserValidatePassword(t *testing.T) {
	user, err := NewUser("Mateus Dias", "maugusto.diaas@gmail.com", "123456", "37386390041", true)
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("1234567"))
}
