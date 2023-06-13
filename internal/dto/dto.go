package dto

import "github.com/mdiaas/goapi/pkg/entity"

type CreateGymClassInput struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

type UpdateGymClassInput struct {
	ID   entity.ID `json:"-"`
	Name string    `json:"name"`
	Link string    `json:"link"`
}

type CreateUserInput struct {
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	CPF      string `json:"cpf"`
	IsAdmin  bool   `json:"is_admin"`
}
