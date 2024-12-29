package main

import (
	"github.com/victorfr4nca/go-crud/internal/http"
	task_handlers "github.com/victorfr4nca/go-crud/internal/http/handlers/task"
	task_service "github.com/victorfr4nca/go-crud/internal/service/task"
	task_storage "github.com/victorfr4nca/go-crud/internal/storage/inmemory/task"
)

func main() {
	taskStorage := task_storage.NewInMemoryTasksStorage()
	taskService := task_service.NewTaskService(taskStorage)
	taskHandler := task_handlers.NewTaskHandlers(taskService)

	server := http.NewServer(taskHandler)
	server.WithPort(":3000")
	if err := server.Start(); err != nil {
		panic(err)
	}
}
