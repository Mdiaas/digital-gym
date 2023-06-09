package database

import (
	"github.com/mdiaas/goapi/internal/entity"
	"gorm.io/gorm"
)

type GymClass struct {
	DB *gorm.DB
}

func NewGymClass(db *gorm.DB) *GymClass {
	return &GymClass{DB: db}
}

func (g *GymClass) Create(gymClass *entity.GymClass) error {
	return g.DB.Create(gymClass).Error
}
func (g *GymClass) FindByID(id string) (*entity.GymClass, error) {
	var gymClass entity.GymClass
	err := g.DB.First(&gymClass, "id = ?", id).Error
	return &gymClass, err
}
func (g *GymClass) Update(gymClass *entity.GymClass) error {
	_, err := g.FindByID(gymClass.ID.String())
	if err != nil {
		return err
	}
	return g.DB.Save(gymClass).Error
}
func (g *GymClass) Delete(gymClass *entity.GymClass) error {
	gymClass, err := g.FindByID(gymClass.ID.String())
	if err != nil {
		return err
	}
	return g.DB.Delete(gymClass).Error
}
func (g *GymClass) FindAll(page, limit int, sort string) ([]entity.GymClass, error) {
	var gymClasses []entity.GymClass
	var err error
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		err = g.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&gymClasses).Error
	} else {
		err = g.DB.Order("created_at " + sort).Find(&gymClasses).Error
	}
	if err != nil {
		return nil, err
	}
	return gymClasses, nil
}
