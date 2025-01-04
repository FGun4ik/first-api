package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type requestBody struct {
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var body requestBody

	json.NewDecoder(r.Body).Decode(&body)
	message := Message{Task: body.Task, IsDone: body.IsDone}
	DB.Create(&message)

	fmt.Fprintf(w, "Message saved: %s", message.Task)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	var messageArr []Message
	DB.Find(&messageArr)

	json.NewEncoder(w).Encode(messageArr)
}

func main() {
	InitDB()

	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/hello", PostHandler).Methods("POST")
	router.HandleFunc("/api/hello", GetHandler).Methods("GET")
	http.ListenAndServe(":8080", router)
}
