package main

import (
	"FirstProject/internal/database"
	"FirstProject/internal/handlers"
	"FirstProject/internal/taskService"
	"FirstProject/internal/userService"
	"FirstProject/internal/web/tasks"
	"FirstProject/internal/web/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	database.InitDB()
	if err := database.DB.AutoMigrate(&taskService.Task{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	if err := database.DB.AutoMigrate(&userService.User{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	tasksRepo := taskService.NewTaskRepository(database.DB)
	tasksService := taskService.NewService(tasksRepo)
	tasksHandler := handlers.NewHandler(tasksService)

	usersRepo := userService.NewUserRepository(database.DB)
	usersService := userService.NewUserService(usersRepo)
	usersHandler := handlers.NewUserHandler(usersService)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(tasksHandler, nil)
	strictUserHandler := users.NewStrictHandler(usersHandler, nil)
	tasks.RegisterHandlers(e, strictHandler)
	users.RegisterHandlers(e, strictUserHandler)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
