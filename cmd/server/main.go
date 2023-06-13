package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mdiaas/goapi/configs"
	"github.com/mdiaas/goapi/internal/entity"
	"github.com/mdiaas/goapi/internal/infra/database"
	"github.com/mdiaas/goapi/internal/infra/webserver/handlers"
	"github.com/mdiaas/goapi/internal/usecases"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.GymClass{}, &entity.User{})
	gymClassDB := database.NewGymClass(db)
	userDB := database.NewUser(db)
	createGymClassUC := usecases.NewCreateGymClassUC(gymClassDB)
	getGymClassUC := usecases.NewGetGymClassUC(gymClassDB)
	updateGymClassUC := usecases.NewUpdateGymClassUC(gymClassDB)
	deleteGymClassUC := usecases.NewDeleteGymClassUC(gymClassDB)
	findAllGymClassesUC := usecases.NewFindAllGymClassesUC(gymClassDB)
	gymClassHandler := handlers.NewGymClassHandler(createGymClassUC, getGymClassUC, updateGymClassUC, deleteGymClassUC, findAllGymClassesUC)

	createUserUC := usecases.NewCreateUserUC(userDB)
	userHandler := handlers.NewUserHandler(createUserUC)
	r := chi.NewRouter()
	r.Post("/gymclass", gymClassHandler.CreateGymClass)
	r.Get("/gymclass", gymClassHandler.FindAllGymClasses)
	r.Get("/gymclass/{id}", gymClassHandler.GetGymClass)
	r.Put("/gymclass/{id}", gymClassHandler.UpdateGymClass)
	r.Delete("/gymclass/{id}", gymClassHandler.DeleteGymClass)

	r.Post("/user", userHandler.CreateUser)
	http.ListenAndServe(":8080", r)

}
