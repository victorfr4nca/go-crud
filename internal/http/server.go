package http

import (
	"net/http"

	task_handlers "github.com/victorfr4nca/go-crud/internal/http/handlers/task"
)

type Server struct {
	port         string
	taskHandler *task_handlers.TaskHandler
}

func NewServer(taskHandler *task_handlers.TaskHandler) *Server {
	return &Server{
		port:         ":3000",
		taskHandler: taskHandler,
	}
}

func (s *Server) WithPort(port string) *Server {
	s.port = port
	return s
}

func (s *Server) Start() error {
	router := http.NewServeMux()

	s.taskHandler.RegisterRoutes(router)

	middlewareChain := MiddlewareChain(
		LoggingMiddleware,
	)

	server := http.Server{
		Addr:    s.port,
		Handler: middlewareChain(router),
	}

	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
