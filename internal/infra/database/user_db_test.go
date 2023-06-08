package database

import (
	"testing"

	"github.com/mdiaas/goapi/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})
	user, err := entity.NewUser("Mateus Dias", "maugusto.diaas@gmail.com", "123456", "37386390041", true)
	if err != nil {
		t.Error(err)
	}
	userDB := NewUser(db)
	err = userDB.Create(user)
	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.FullName, userFound.FullName)
	assert.Equal(t, user.Email, userFound.Email)
	assert.Equal(t, user.CPF, userFound.CPF)
	assert.Equal(t, user.IsAdmin, userFound.IsAdmin)
}

func TestFindByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("Mateus Dias", "maugusto.diaas@gmail.com", "123456", "37386390041", true)
	userDB := NewUser(db)
	err = userDB.Create(user)
	assert.Nil(t, err)

	userFound, err := userDB.FindByID(user.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.FullName, userFound.FullName)
	assert.Equal(t, user.Email, userFound.Email)
	assert.Equal(t, user.CPF, userFound.CPF)
	assert.Equal(t, user.IsAdmin, userFound.IsAdmin)
}
