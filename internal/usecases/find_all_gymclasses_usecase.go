package usecases

import (
	"github.com/mdiaas/goapi/internal/entity"
	"github.com/mdiaas/goapi/internal/infra/database"
)

type FindAllGymClassesUC struct {
	GymClassDB database.GymClassDatabaseInterface
}

func NewFindAllGymClassesUC(db database.GymClassDatabaseInterface) *FindAllGymClassesUC {
	return &FindAllGymClassesUC{
		GymClassDB: db,
	}
}

func (h *FindAllGymClassesUC) Execute(page, limit int, sort string) ([]entity.GymClass, error) {

	gymClasses, err := h.GymClassDB.FindAll(page, limit, sort)
	if err != nil {
		return nil, err
	}
	return gymClasses, nil
}
