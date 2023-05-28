package entity

import (
	"github.com/mdiaas/goapi/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       entity.ID  `json:"id"`
	FullName string     `json:"fullname"`
	Email    string     `json:"email"`
	Password string     `json:"-"`
	CPF      entity.CPF `json:"cpf"`
	IsAdmin  bool       `json:"is_admin" `
}

func NewUser(fullname, email, password, cpfString string, isAdmin bool) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	cpf, err := entity.NewCPF(cpfString)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       entity.NewID(),
		FullName: fullname,
		Email:    email,
		Password: string(hash),
		CPF:      cpf,
		IsAdmin:  isAdmin,
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
