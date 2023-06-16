package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/mdiaas/goapi/internal/dto"
	"github.com/mdiaas/goapi/internal/usecases"
)

type GymClassHandler struct {
	CreateGymClassUC    usecases.CreateGymClassUCInterface
	GetGymClassUC       usecases.GetGymClassUCInterface
	UpdateGymClassUC    usecases.UpdateGymClassUCInterface
	DeleteGymClassUC    usecases.DeleteGymClassUCInterface
	FindAllGymClassesUC usecases.FindAllGymClassesUCInterface
}

func NewGymClassHandler(createGymClassUC usecases.CreateGymClassUCInterface, getGymClassUC usecases.GetGymClassUCInterface, updateGymClassUC usecases.UpdateGymClassUCInterface, deleteGymClassUC usecases.DeleteGymClassUCInterface, findAllGymClassesUC usecases.FindAllGymClassesUCInterface) *GymClassHandler {
	return &GymClassHandler{
		CreateGymClassUC:    createGymClassUC,
		GetGymClassUC:       getGymClassUC,
		UpdateGymClassUC:    updateGymClassUC,
		DeleteGymClassUC:    deleteGymClassUC,
		FindAllGymClassesUC: findAllGymClassesUC,
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

func (h *GymClassHandler) DeleteGymClass(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	gymClass, err := h.GetGymClassUC.Execute(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = h.DeleteGymClassUC.Execute(gymClass)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func (h *GymClassHandler) FindAllGymClasses(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}
	sort := r.URL.Query().Get("sort")
	if sort != "desc" {
		sort = "asc"
	}

	gymClasses, err := h.FindAllGymClassesUC.Execute(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(gymClasses)
}
