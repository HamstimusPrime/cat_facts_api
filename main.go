package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

type  catAPIresult struct{
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

// Middleware that wraps the handler with apiURL
func fetchCatFactsMiddleware(apiURL string, handler func(string, http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(apiURL, w, r)
	}
}

func fetchCatFacts(apiURL string, w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(apiURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode > 200 {
		fmt.Printf("request error! status code: %v\n", resp.StatusCode)
		return
	}

	//parse response body from JSON to GO Struct
	fmt.Print("request made to server\n")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//find a way to pass the apiURL to the function
	apiURL := os.Getenv("API_URL")
	port := os.Getenv("PORT")
	mux := http.NewServeMux()
	mux.HandleFunc("GET /me", fetchCatFactsMiddleware(apiURL, fetchCatFacts))

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("server running on port: %v\n", port)
	log.Fatal(server.ListenAndServe())
}



// {
//   "status": "success",
//   "user": {
//     "email": "<your email>",
//     "name": "<your full name>",
//     "stack": "<your backend stack>"
//   },
//   "timestamp": "<current UTC time in ISO 8601 format>",
//   "fact": "<random cat fact from Cat Facts API>"
// }
