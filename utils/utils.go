package utils

import (
	"encoding/json"
	"log"
	"net/http"
)



type errJSON struct{
	err string
	status int
}

func RespondWithError(writer http.ResponseWriter, HTTPstatus int){
	errMessage := "unable to compelete request"
	errJSON, err := json.Marshal(errJSON{err: errMessage,
										status: HTTPstatus,})
	if err != nil{
		log.Fatal("unable to parse response JSON")
	}
	writer.WriteHeader(HTTPstatus)
	writer.Write([]byte(errJSON))
}