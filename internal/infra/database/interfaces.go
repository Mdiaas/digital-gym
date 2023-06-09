package database

import "github.com/mdiaas/goapi/internal/entity"

type UserDatabaseInterface interface {
	Create(user *entity.User) error
	FindByID(id string) (*entity.User, error)
}

type GymClassDatabaseInterface interface {
	Create(gymClass *entity.GymClass) error
	FindAll(page, limit int, sort string) ([]entity.GymClass, error)
	FindByID(id string) (*entity.GymClass, error)
	Update(gymClass *entity.GymClass) error
	Delete(gymClass *entity.GymClass) error
}
