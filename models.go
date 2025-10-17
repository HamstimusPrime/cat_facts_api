package main

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

type metadata struct {
	ApiURL string
	Email  string
	Name   string
	Stack  string
	Timeout string
	Status string
}
