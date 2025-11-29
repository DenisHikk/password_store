package server

import (
	auth "genpasstore/internal/auth/app"
	authUser "genpasstore/internal/auth/handler"
	password "genpasstore/internal/password/handler"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Deps struct {
	AuthHandler  *authUser.AuthHandler
	TokenManager *auth.TokenManager
}

func NewHTTPServer(deps Deps) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("{\"status\": \"ok\"}"))
	})

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/register", deps.AuthHandler.HandleRegistry)
			r.Post("/login", deps.AuthHandler.HandleLogin)
		})
		r.Group(func(r chi.Router) {
			r.Use(deps.TokenManager.MiddlewareJWTToken)
			r.Post("/refresh", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("{\"TODO\": \"done in future\"}"))
			})
			r.Route("/utils", func(r chi.Router) {
				r.Post("/password-generator", password.HandleGeneratePassword)
			})
		})

	})

	log.Println("Done registry handler")
	return r
}
