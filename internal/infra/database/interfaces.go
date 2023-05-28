package database

import "github.com/mdiaas/goapi/internal/entity"

type UserDatabaseInterface interface {
	Create(user *entity.User) error
	FindByID(id string) (*entity.User, error)
}
