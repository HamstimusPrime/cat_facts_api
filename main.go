package main

import (
	"encoding/json"
	"time"
	"log"
	"net/http"
	"os"

	"github.com/HamstimusPrime/cat_facts_api/utils"
	"github.com/joho/godotenv"
)

type response struct {
	Status string `json:"status"`
	User   struct {
		Email string `json:"email"`
		Name  string `json:"name"`
		Stack string `json:"stack"`
	} `json:"user"`
	Timestamp string `json:"timestamp"`
	Fact      string `json:"fact"`
}

type catAPIresult struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

type metadata struct{
	ApiURL string
	Email string
	Name string
	Stack string
	Status string
}

// Middleware wraps the fetchCatFacts handler function
func fetchCatFactsMiddleware(metadata metadata, handler func(metadata, http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(metadata, w, r)
	}
}

func fetchCatFacts(metadata metadata, w http.ResponseWriter, r *http.Request) {
	client := &http.Client{Timeout: 5 * time.Second}
	
	resp, err := client.Get(metadata.ApiURL)
	if err != nil {
		utils.RespondWithError(w, http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	var catFact catAPIresult
	if err := json.NewDecoder(resp.Body).Decode(&catFact); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError)
		return
	}

	response := response{
		Status: metadata.Status,
		User: struct {
			Email string `json:"email"`
			Name  string `json:"name"`
			Stack string `json:"stack"`
		}{
			Email: metadata.Email,
			Name:  metadata.Name,
			Stack: metadata.Stack,
		},
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Fact:     catFact.Fact,
	}

	respondWithJSON(response,w,200)
}

func respondWithJSON(responseObject interface{}, w http.ResponseWriter, HTTPstatus int){
	respJSON, err := json.Marshal(responseObject)
	if err != nil {
		utils.RespondWithError(w, http.StatusServiceUnavailable)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(HTTPstatus)
	w.Write([]byte(respJSON))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	metadata := metadata{
		ApiURL: os.Getenv("API_URL"),
		Email: os.Getenv("EMAIL"),
		Name: os.Getenv("NAME"),
		Stack: os.Getenv("STACK"),
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
