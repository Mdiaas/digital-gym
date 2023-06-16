package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/mdiaas/goapi/internal/dto"
	"github.com/mdiaas/goapi/internal/usecases"
)

type UserHandler struct {
	CreateUserUC usecases.CreateUserUCInterface
	LoginUserUC  usecases.LoginUserUCInterface
}

func NewUserHandler(createUserUC usecases.CreateUserUCInterface, loginUserUC usecases.LoginUserUCInterface) *UserHandler {
	return &UserHandler{
		CreateUserUC: createUserUC,
		LoginUserUC:  loginUserUC,
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
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	expiresIn := r.Context().Value("expiresIn").(int)
	var userDto *dto.GetJWTUserInput
	err := json.NewDecoder(r.Body).Decode(&userDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err := h.LoginUserUC.Execute(userDto)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	_, token, err := jwt.Encode(map[string]interface{}{
		"sub":   user.ID.String(),
		"exp":   time.Now().Add(time.Second * time.Duration(expiresIn)).Unix(),
		"admin": user.IsAdmin,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	accessToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: token,
	}
	w.Header().Set("Content-type", "Application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}
