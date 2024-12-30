package http

import (
	"net/http"

	"github.com/victorfr4nca/go-crud/internal/http/task"
)

type Server struct {
	port        string
	taskHandler *task.Handler
}

func NewServer(taskHandler *task.Handler) *Server {
	return &Server{
		port:        ":3000",
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
