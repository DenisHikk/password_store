package main

import (
	"context"
	"genpasstore/internal/db"
	"genpasstore/internal/server"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := godotenv.Load()
	if err != nil {
		return err
	}

	configDB, err := loadConfigDB()
	if err != nil {
		return err
	}
	pool, err := db.NewPool(ctx, configDB)
	if err != nil {
		return err
	}
	defer pool.Close()

	srv := server.NewHTTPServer(pool)

	log.Printf("Server start and listen in {}:8000")
	err = http.ListenAndServe(":8000", srv)
	if err != nil {
		return err
	}
	return nil
}
