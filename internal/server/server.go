package server

import (
	password "genpasstore/internal/password/handler"
	authUser "genpasstore/internal/user/handler"
	userRepository "genpasstore/internal/user/repository"
	userService "genpasstore/internal/user/service"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewHTTPServer(pool *pgxpool.Pool) *http.ServeMux {
	mux := http.NewServeMux()

	userRepo := userRepository.NewUserRepository(pool)
	userService := userService.NewUserService(userRepo)
	authHandler := authUser.NewAuthHandler(userService)

	// handler password
	mux.HandleFunc("/password/generate", password.HandleGeneratePassword)
	// handler user
	mux.HandleFunc("/user/registry", authHandler.HandleRegistry)
	mux.HandleFunc("/user/login", authHandler.HandleLogin)

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("{\"status\": \"ok\"}")) })

	log.Println("Done registry handler")
	return mux
}
