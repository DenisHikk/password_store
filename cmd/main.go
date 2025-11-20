package main

import (
	"context"
	authUser "genpasstore/internal/auth/handler"
	"genpasstore/internal/db"
	"genpasstore/internal/httpx/server"
	userRepository "genpasstore/internal/user/repository"
	userService "genpasstore/internal/user/service"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"

	_ "net/http/pprof"
)

func loadConfigDB() (db.ConfigDB, error) {
	PostgresUser := os.Getenv("POSTGRES_USER")
	PostgresPassword := os.Getenv("POSTGRES_PASSWORD")
	PostgresDB := os.Getenv("POSTGRES_DB")
	PostgresHost := os.Getenv("POSTGRES_HOST")
	PostgresPort, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		return db.ConfigDB{}, err
	}
	configPostgres := db.ConfigDB{
		User:              PostgresUser,
		Password:          PostgresPassword,
		DBName:            PostgresDB,
		Port:              PostgresPort,
		Host:              PostgresHost,
		ConnectionTimeout: 5 * time.Second,
	}

	return configPostgres, nil
}

func main() {
	err := realMain()
	if err != nil {
		log.Fatal(err)
	}
}

func realMain() error {
	ctx := context.Background()
	log.Println("Load dotenv")
	err := godotenv.Load()
	if err != nil {
		return err
	}

	log.Println("Load config for DB")
	configDB, err := loadConfigDB()
	if err != nil {
		return err
	}

	log.Println("Create pool for DB")
	pool, err := db.NewPool(ctx, configDB)
	if err != nil {
		return err
	}
	defer pool.Close()

	userRepo := userRepository.NewUserRepository(pool)
	userService := userService.NewUserService(userRepo)
	authHandler := authUser.NewAuthHandler(userService)

	log.Println("Starting server...")
	srv := server.NewHTTPServer(server.Deps{
		AuthHandler: authHandler,
	})

	log.Printf("Server start and listen in 0.0.0.0:8000")
	err = http.ListenAndServe("0.0.0.0:8000", srv)
	if err != nil {
		return err
	}
	return nil
}
