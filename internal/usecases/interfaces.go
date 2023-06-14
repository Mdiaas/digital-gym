package usecases

import (
	"io"

	"github.com/mdiaas/goapi/internal/dto"
	"github.com/mdiaas/goapi/internal/entity"
)

type CreateGymClassUCInterface interface {
	Execute(gymClassDto *dto.CreateGymClassInput) error
}
type GetGymClassUCInterface interface {
	Execute(id string) (*entity.GymClass, error)
}
type UpdateGymClassUCInterface interface {
	Execute(id string, requestBody io.ReadCloser) (*entity.GymClass, error)
}
type DeleteGymClassUCInterface interface {
	Execute(gymClass *entity.GymClass) error
}
type FindAllGymClassesUCInterface interface {
	Execute(page, limit int, sort string) ([]entity.GymClass, error)
}
type CreateUserUCInterface interface {
	Execute(user *dto.CreateUserInput) error
}
type LoginUserUCInterface interface {
	Execute(userDto dto.GetJWTUserInput) (*entity.User, error)
}
