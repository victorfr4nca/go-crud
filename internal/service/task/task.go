package task

import (
	"github.com/victorfr4nca/go-crud/internal/service"
	"github.com/victorfr4nca/go-crud/internal/storage"
	"github.com/victorfr4nca/go-crud/internal/types"
)

var _ service.TaskService = (*TaskService)(nil)

type TaskService struct {
	storage storage.TasksStorage
}

func NewTaskService(storage storage.TasksStorage) *TaskService {
	return &TaskService{
		storage: storage,
	}
}

func (ts *TaskService) List() ([]*types.Task, error) {
	return ts.storage.List()
}

func (ts *TaskService) Create(task *types.Task) (*types.Task, error) {
	err := ts.storage.Save(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (ts *TaskService) Update(task *types.Task) (*types.Task, error) {
	err := ts.storage.Update(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (ts *TaskService) Delete(id string) error {
	return ts.storage.Delete(id)
}
