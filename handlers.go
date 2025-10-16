package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/HamstimusPrime/cat_facts_api/utils"
)
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
		Fact:      catFact.Fact,
	}

	respondWithJSON(response, w, 200)
}

func respondWithJSON(responseObject interface{}, w http.ResponseWriter, HTTPstatus int) {
	respJSON, err := json.Marshal(responseObject)
	if err != nil {
		utils.RespondWithError(w, http.StatusServiceUnavailable)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(HTTPstatus)
	w.Write([]byte(respJSON))
}
