package task

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/victorfr4nca/go-crud/internal/entity"
	"github.com/victorfr4nca/go-crud/internal/service/task"
)

type Handler struct {
	taskService task.Service
}

func New(taskService task.Service) *Handler {
	return &Handler{
		taskService: taskService,
	}
}

func (th *Handler) GetHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := th.taskService.List()
	if err != nil {
		http.Error(w, "Failed to list tasks", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tasks)
}

func (th *Handler) PostHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	task := &entity.Task{}

	err = json.Unmarshal(body, &task)
	if err != nil {
		http.Error(w, "Failed to unmarshal request body", http.StatusInternalServerError)
		return
	}

	_, err = th.taskService.Create(task)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create task: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(task)
}

func (th *Handler) PatchHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	intId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Failed to parse id", http.StatusInternalServerError)
		return
	}

	task := &entity.Task{
		Id: intId,
	}

	err = json.Unmarshal(body, &task)
	if err != nil {
		http.Error(w, "Failed to unmarshal request body", http.StatusInternalServerError)
		return
	}

	_, err = th.taskService.Update(task)
	if err != nil {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(task)
}

func (th *Handler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	err := th.taskService.Delete(id)
	if err != nil {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}
}
