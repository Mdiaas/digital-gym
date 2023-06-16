package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/mdiaas/goapi/configs"
	"github.com/mdiaas/goapi/internal/entity"
	"github.com/mdiaas/goapi/internal/infra/database"
	"github.com/mdiaas/goapi/internal/infra/webserver/handlers"
	"github.com/mdiaas/goapi/internal/infra/webserver/middlewares"
	"github.com/mdiaas/goapi/internal/usecases"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	configs, err := configs.LoadConfig(".")
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
	loginUserUC := usecases.NewLoginUC(userDB)
	userHandler := handlers.NewUserHandler(createUserUC, loginUserUC)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("expiresIn", configs.JWTExpiresIn))
	r.Route("/gymclass", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Get("/", gymClassHandler.FindAllGymClasses)
		r.Get("/{id}", gymClassHandler.GetGymClass)
		r.Group(func(r chi.Router) {
			r.Use(middlewares.IsAdminMiddleware)
			r.Post("/", gymClassHandler.CreateGymClass)
			r.Put("/{id}", gymClassHandler.UpdateGymClass)
			r.Delete("/{id}", gymClassHandler.DeleteGymClass)
		})
	})
	r.Route("/user", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Use(middlewares.IsAdminMiddleware)
		r.Post("/", userHandler.CreateUser)
	})
	r.Post("/user/login", userHandler.GetJWT)
	http.ListenAndServe(":8080", r)
}
