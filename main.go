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

	metadata := metadata{
		ApiURL: os.Getenv("API_URL"),
		Email:  os.Getenv("EMAIL"),
		Name:   os.Getenv("NAME"),
		Stack:  os.Getenv("STACK"),
		Status: os.Getenv("STATUS"),
	}

	port := os.Getenv("PORT")
	mux := http.NewServeMux()
	mux.HandleFunc("GET /me", fetchCatFactsMiddleware(metadata, fetchCatFacts))

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("server running on port: %v\n", port)
	log.Fatal(server.ListenAndServe())
}
