package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mdiaas/goapi/internal/dto"
	"github.com/mdiaas/goapi/internal/entity"
	"github.com/mdiaas/goapi/internal/infra/database"
	entityPkg "github.com/mdiaas/goapi/pkg/entity"
)

type GymClassHandler struct {
	GymClassDB database.GymClassDatabaseInterface
}

func NewGymClassHandler(db database.GymClassDatabaseInterface) *GymClassHandler {
	return &GymClassHandler{
		GymClassDB: db,
	}
}

func (h *GymClassHandler) CreateGymClass(w http.ResponseWriter, r *http.Request) {
	var gymClass dto.CreateGymClassInput
	err := json.NewDecoder(r.Body).Decode(&gymClass)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	g, err := entity.NewGymClass(gymClass.Name, gymClass.Link)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.GymClassDB.Create(g)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
func (h *GymClassHandler) GetGymClass(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	gymClass, err := h.GymClassDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(gymClass)
}
func (h *GymClassHandler) UpdateGymClass(w http.ResponseWriter, r *http.Request) {
	var gymClass entity.GymClass
	id := chi.URLParam(r, "id")
	_, err := h.GymClassDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	err = json.NewDecoder(r.Body).Decode(&gymClass)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	gymClass.ID, err = entityPkg.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	err = h.GymClassDB.Update(&gymClass)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
