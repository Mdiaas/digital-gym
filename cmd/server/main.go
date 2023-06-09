package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mdiaas/goapi/configs"
	"github.com/mdiaas/goapi/internal/entity"
	"github.com/mdiaas/goapi/internal/infra/database"
	"github.com/mdiaas/goapi/internal/infra/webserver/handlers"
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
	gymClassHandler := handlers.NewGymClassHandler(gymClassDB)
	r := chi.NewRouter()
	r.Post("/gymclass", gymClassHandler.CreateGymClass)
	r.Get("/gymclass/{id}", gymClassHandler.GetGymClass)
	r.Put("/gymclass/{id}", gymClassHandler.UpdateGymClass)
	http.ListenAndServe(":8080", r)

}
