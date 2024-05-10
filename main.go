package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	InitRoutes()

	port := os.Getenv("PORT")
	log.Printf("Server is starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
