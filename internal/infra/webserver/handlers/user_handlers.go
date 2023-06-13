package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mdiaas/goapi/internal/dto"
	"github.com/mdiaas/goapi/internal/usecases"
)

type UserHandler struct {
	CreateUserUC usecases.CreateUserUCInterface
}

func NewUserHandler(createUserUC usecases.CreateUserUCInterface) *UserHandler {
	return &UserHandler{
		CreateUserUC: createUserUC,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user *dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.CreateUserUC.Execute(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
}
