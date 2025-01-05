package main

import (
	"FirstProject/internal/database"
	"FirstProject/internal/handlers"
	"FirstProject/internal/taskService"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)

	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/get", handler.GetTasksHandler).Methods("GET")
	router.HandleFunc("/api/post", handler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/delete/{id:[0-9]+}", handler.DeleteTaskHandler).Methods("DELETE")
	router.HandleFunc("/api/update/{id:[0-9]+}", handler.UpdateTaskHandler).Methods("PUT", "PATCH")
	http.ListenAndServe(":8080", router)
}
