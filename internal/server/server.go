package server

import (
	password "genpasstore/internal/password/handler"
	authUser "genpasstore/internal/user/handler"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Deps struct {
	AuthHandler *authUser.AuthHandler
}

func NewHTTPServer(deps Deps) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("{\"status\": \"ok\"}"))
	})

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/password", func(r chi.Router) {
			r.Post("/generate", password.HandleGeneratePassword)
		})
		r.Route("/auth", func(r chi.Router) {
			r.Post("/registry", deps.AuthHandler.HandleRegistry)
			r.Post("/login", deps.AuthHandler.HandleLogin)
		})
	})

	log.Println("Done registry handler")
	return r
}
