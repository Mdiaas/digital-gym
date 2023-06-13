package usecases

import (
	"github.com/mdiaas/goapi/internal/dto"
	"github.com/mdiaas/goapi/internal/entity"
	"github.com/mdiaas/goapi/internal/infra/database"
)

type CreateUserUC struct {
	UserDB database.UserDatabaseInterface
}

func NewCreateUserUC(db database.UserDatabaseInterface) *CreateUserUC {
	return &CreateUserUC{
		UserDB: db,
	}
}

func (c *CreateUserUC) Execute(userDto *dto.CreateUserInput) error {
	user, err := entity.NewUser(userDto.FullName, userDto.Email, userDto.Password, userDto.CPF, userDto.IsAdmin)
	if err != nil {
		return err
	}
	err = c.UserDB.Create(user)
	if err != nil {
		return err
	}
	return nil
}
