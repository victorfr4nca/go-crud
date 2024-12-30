package main

import (
	"github.com/victorfr4nca/go-crud/internal/database"
	"github.com/victorfr4nca/go-crud/internal/http"
	task_handler "github.com/victorfr4nca/go-crud/internal/http/task"
	"github.com/victorfr4nca/go-crud/internal/repository/task/sqlite"
	"github.com/victorfr4nca/go-crud/internal/service/task"
)

func main() {
	// inMemoryTaskRepository := memory.New()
	db, err := database.Init("crud.db")
	if err != nil {
		panic(err)
	}

	taskRepository := sqlite.New(db)
	taskService := task.New(taskRepository)
	taskHandler := task_handler.New(taskService)

	server := http.NewServer(taskHandler)
	server.WithPort(":3000")
	if err := server.Start(); err != nil {
		panic(err)
	}
}
