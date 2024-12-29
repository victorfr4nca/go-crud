package task

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/victorfr4nca/go-crud/internal/service"
	"github.com/victorfr4nca/go-crud/internal/types"
)

type TaskHandler struct {
	taskService service.TaskService
}

func NewTaskHandlers(taskService service.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

func (th *TaskHandler) GetHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := th.taskService.List()
	if err != nil {
		http.Error(w, "Failed to list tasks", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tasks)
}

func (th *TaskHandler) PostHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	task := &types.Task{}

	err = json.Unmarshal(body, &task)
	if err != nil {
		http.Error(w, "Failed to unmarshal request body", http.StatusInternalServerError)
		return
	}

	_, err = th.taskService.Create(task)
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(task)
}

func (th *TaskHandler) PatchHandler(w http.ResponseWriter, r *http.Request) {
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

	task := &types.Task{
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

func (th *TaskHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	err := th.taskService.Delete(id)
	if err != nil {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}
}
