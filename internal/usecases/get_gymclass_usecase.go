package usecases

import (
	"github.com/mdiaas/goapi/internal/entity"
	"github.com/mdiaas/goapi/internal/infra/database"
)

type GetGymClassUC struct {
	GymClassDB database.GymClassDatabaseInterface
}

func NewGetGymClassUC(db database.GymClassDatabaseInterface) *GetGymClassUC {
	return &GetGymClassUC{
		GymClassDB: db,
	}
}

func (c *GetGymClassUC) Execute(id string) (*entity.GymClass, error) {
	gymClass, err := c.GymClassDB.FindByID(id)
	if err != nil {
		return nil, err
	}
	return gymClass, nil
}
