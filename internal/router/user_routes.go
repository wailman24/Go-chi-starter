package router

import (
	"github.com/go-chi/chi"
	"github.com/wailman24/Go-chi-starter.git/internal/handlers"
	"github.com/wailman24/Go-chi-starter.git/internal/middlewares"
)

func UserRoutes() *chi.Mux {
	r := chi.NewRouter()
	h := handlers.NewUserHandler()
	r.Post("/login", h.Login)
	//r.Post("/create", h.Register)

	r.Group(func(r chi.Router) {
		r.Use(middlewares.AuthMiddleware)
		r.Post("/create", h.Register)
	})

	return r
}
