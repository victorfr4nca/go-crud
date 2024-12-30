package task

import (
	"github.com/victorfr4nca/go-crud/internal/entity"
	"github.com/victorfr4nca/go-crud/internal/repository/task"
)

type Service interface {
	List() ([]*entity.Task, error)
	Create(task *entity.Task) (*entity.Task, error)
	Update(task *entity.Task) (*entity.Task, error)
	Delete(id string) error
}


type taskService struct {
	taskRepository task.Repository
}

var _ Service = (*taskService)(nil)

func New(taskRepository task.Repository) *taskService {
	return &taskService{
		taskRepository: taskRepository,
	}
}

func (ts *taskService) List() ([]*entity.Task, error) {
	return ts.taskRepository.List()
}

func (ts *taskService) Create(task *entity.Task) (*entity.Task, error) {
	err := ts.taskRepository.Save(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (ts *taskService) Update(task *entity.Task) (*entity.Task, error) {
	err := ts.taskRepository.Update(task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (ts *taskService) Delete(id string) error {
	return ts.taskRepository.Delete(id)
}
