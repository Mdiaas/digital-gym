package usecases

import (
	"github.com/mdiaas/goapi/internal/entity"
	"github.com/mdiaas/goapi/internal/infra/database"
)

type DeleteGymClassUC struct {
	GymClassDB database.GymClassDatabaseInterface
}

func NewDeleteGymClassUC(db database.GymClassDatabaseInterface) *DeleteGymClassUC {
	return &DeleteGymClassUC{
		GymClassDB: db,
	}
}

func (h *DeleteGymClassUC) Execute(gymClass *entity.GymClass) error {
	h.GymClassDB.Delete(gymClass)
	return nil
}
