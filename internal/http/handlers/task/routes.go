package task

import "net/http"

func (th *TaskHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /tasks", th.GetHandler)
	router.HandleFunc("POST /tasks", th.PostHandler)
	router.HandleFunc("PATCH /tasks/{id}", th.PatchHandler)
	router.HandleFunc("DELETE /tasks/{id}", th.DeleteHandler)
}
