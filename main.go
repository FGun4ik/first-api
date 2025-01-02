package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type requestBody struct {
	Message string `json:"message"`
}

var task string

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	var body requestBody

	json.NewDecoder(r.Body).Decode(&body)
	task = body.Message

	fmt.Fprintf(w, "Message saved: %s", task)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s", task)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", HelloHandler).Methods("POST")
	router.HandleFunc("/api/hello", GetHandler).Methods("GET")
	http.ListenAndServe(":8080", router)
}
