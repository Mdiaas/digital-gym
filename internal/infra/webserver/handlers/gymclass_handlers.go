package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mdiaas/goapi/internal/dto"
	"github.com/mdiaas/goapi/internal/usecases"
)

type GymClassHandler struct {
	CreateGymClassUC usecases.CreateGymClassUCInterface
	GetGymClassUC    usecases.GetGymClassUCInterface
	UpdateGymClassUC usecases.UpdateGymClassUCInterface
}

func NewGymClassHandler(createGymClassUC usecases.CreateGymClassUCInterface, getGymClassUC usecases.GetGymClassUCInterface, updateGymClassUC usecases.UpdateGymClassUCInterface) *GymClassHandler {
	return &GymClassHandler{
		CreateGymClassUC: createGymClassUC,
		GetGymClassUC:    getGymClassUC,
		UpdateGymClassUC: updateGymClassUC,
	}
}

func (h *GymClassHandler) CreateGymClass(w http.ResponseWriter, r *http.Request) {
	var gymClass *dto.CreateGymClassInput
	err := json.NewDecoder(r.Body).Decode(&gymClass)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.CreateGymClassUC.Execute(gymClass)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *GymClassHandler) GetGymClass(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	gymClass, err := h.GetGymClassUC.Execute(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(gymClass)
}

func (h *GymClassHandler) UpdateGymClass(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	_, err := h.GetGymClassUC.Execute(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	gymClass, err := h.UpdateGymClassUC.Execute(id, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(gymClass)
}
