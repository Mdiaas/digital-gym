package usecases

import (
	"errors"

	"github.com/mdiaas/goapi/internal/dto"
	"github.com/mdiaas/goapi/internal/entity"
	"github.com/mdiaas/goapi/internal/infra/database"
)

type GetJwtUC struct {
	UserDB database.UserDatabaseInterface
}

func NewGetJwtUC(db database.UserDatabaseInterface) *GetJwtUC {
	return &GetJwtUC{
		UserDB: db,
	}
}

func (c *GetJwtUC) Execute(userDto dto.GetJWTUserInput) (*entity.User, error) {
	u, err := c.UserDB.FindByEmail(userDto.Email)
	if err != nil {
		return nil, err
	}
	if u.ValidatePassword(userDto.Password) {
		return nil, errors.New("invalid password")
	}
	return u, nil
}
