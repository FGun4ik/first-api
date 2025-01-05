package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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

	json.NewEncoder(w).Encode(body)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	var messageArr []Message
	DB.Find(&messageArr)

	var result []requestBody
	for _, message := range messageArr {
		result = append(result, requestBody{Task: message.Task, IsDone: message.IsDone})
	}
	json.NewEncoder(w).Encode(result)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var body requestBody
	json.NewDecoder(r.Body).Decode(&body)

	var message Message
	DB.First(&message, id)

	if body.Task != "" {
		message.Task = body.Task
	}
	message.IsDone = body.IsDone
	DB.Save(&message)
	json.NewEncoder(w).Encode(body)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var body requestBody
	json.NewDecoder(r.Body).Decode(&body)

	var message Message
	DB.First(&message, id)
	DB.Delete(&message)
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	InitDB()

	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/hello", PostHandler).Methods("POST")
	router.HandleFunc("/api/hello", GetHandler).Methods("GET")
	router.HandleFunc("/api/hello/{id:[0-9]+}", UpdateHandler).Methods("PUT", "PATCH")
	router.HandleFunc("/api/hello/{id:[0-9]+}", DeleteHandler).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
