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
