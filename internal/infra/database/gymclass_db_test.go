package database

import (
	"fmt"
	"testing"

	"github.com/mdiaas/goapi/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateNewGymClass(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.GymClass{})
	gymClass, err := entity.NewGymClass("yoga", "youtube.com")
	assert.NoError(t, err)
	gymClassDB := NewGymClass(db)
	err = gymClassDB.Create(gymClass)
	assert.NoError(t, err)
	assert.NotEmpty(t, gymClass.ID)
}

func TestFindAllGymClasses(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.GymClass{})

	for i := 0; i < 24; i++ {
		gymClass, err := entity.NewGymClass(fmt.Sprintf("class %d", i+1), "youtube.com")
		assert.NoError(t, err)
		db.Create(gymClass)
	}
	gymClassDB := NewGymClass(db)
	gymClasses, err := gymClassDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, gymClasses, 10)
	assert.Equal(t, gymClasses[0].Name, "class 1")
	assert.Equal(t, gymClasses[9].Name, "class 10")
	gymClasses, err = gymClassDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, gymClasses, 10)
	assert.Equal(t, gymClasses[0].Name, "class 11")
	assert.Equal(t, gymClasses[9].Name, "class 20")
	gymClasses, err = gymClassDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, gymClasses, 4)
	assert.Equal(t, gymClasses[0].Name, "class 21")
	assert.Equal(t, gymClasses[3].Name, "class 24")
}

func TestFindGymClassByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.GymClass{})
	gymClass, err := entity.NewGymClass("Class 1", "youtube.com")
	assert.NoError(t, err)
	db.Create(gymClass)
	gymClassDB := NewGymClass(db)
	gymClassFound, err := gymClassDB.FindByID(gymClass.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, gymClassFound.ID, gymClass.ID)
	assert.Equal(t, gymClassFound.Name, gymClass.Name)
	assert.Equal(t, gymClassFound.Link, gymClass.Link)
}

func TestDeleteGymClass(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.GymClass{})
	gymClass, err := entity.NewGymClass("Class 1", "youtube.com")
	assert.NoError(t, err)
	db.Create(gymClass)
	gymClassDB := NewGymClass(db)
	err = gymClassDB.Delete(gymClass)
	assert.NoError(t, err)

	_, err = gymClassDB.FindByID(gymClass.ID.String())
	assert.Error(t, err)
}
