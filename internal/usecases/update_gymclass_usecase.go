package usecases

import (
	"encoding/json"
	"io"

	"github.com/mdiaas/goapi/internal/entity"
	"github.com/mdiaas/goapi/internal/infra/database"
	entityPkg "github.com/mdiaas/goapi/pkg/entity"
)

type UpdateGymClassUC struct {
	GymClassDB database.GymClassDatabaseInterface
}

func NewUpdateGymClassUC(db database.GymClassDatabaseInterface) *UpdateGymClassUC {
	return &UpdateGymClassUC{
		GymClassDB: db,
	}
}

func (h *UpdateGymClassUC) Execute(id string, requestBody io.ReadCloser) (*entity.GymClass, error) {
	var gymClass entity.GymClass
	err := json.NewDecoder(requestBody).Decode(&gymClass)
	if err != nil {
		return nil, err
	}
	gymClass.ID, err = entityPkg.ParseID(id)
	if err != nil {
		return nil, err
	}
	err = h.GymClassDB.Update(&gymClass)
	if err != nil {
		return nil, err
	}
	return &gymClass, nil
}
