package main

import (
	"genpass/db"
	"genpass/handler"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
		return
	}
	errInitDB := db.InitDB()
	if errInitDB != nil {
		log.Fatalf("Error while connect to DB: %w", errInitDB)
	}

	http.HandleFunc("/password", handler.HandleGeneratePassword)
	http.HandleFunc("/registry", handler.HandleRegistry)

	fs := http.FileServer(http.Dir("web"))
	http.Handle("/", fs)

	log.Println("listen :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))

}
