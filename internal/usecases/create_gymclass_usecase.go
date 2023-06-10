package usecases

import (
	"github.com/mdiaas/goapi/internal/dto"
	"github.com/mdiaas/goapi/internal/entity"
	"github.com/mdiaas/goapi/internal/infra/database"
)

type CreateGymClassUC struct {
	GymClassDB database.GymClassDatabaseInterface
}

func NewCreateGymClassUC(db database.GymClassDatabaseInterface) *CreateGymClassUC {
	return &CreateGymClassUC{
		GymClassDB: db,
	}
}

func (c *CreateGymClassUC) Execute(gymClassDto *dto.CreateGymClassInput) error {
	g, err := entity.NewGymClass(gymClassDto.Name, gymClassDto.Link)
	if err != nil {
		return err
	}
	err = c.GymClassDB.Create(g)
	if err != nil {
		return err
	}
	return nil
}
