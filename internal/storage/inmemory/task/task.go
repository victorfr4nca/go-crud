package task

import (
	"errors"
	"math/rand"
	"strconv"

	"github.com/victorfr4nca/go-crud/internal/storage"
	"github.com/victorfr4nca/go-crud/internal/types"
)

type InMemoryTasksStorage struct {
	tasks []*types.Task
}

var _ storage.TasksStorage = (*InMemoryTasksStorage)(nil)

func NewInMemoryTasksStorage() *InMemoryTasksStorage {
	return &InMemoryTasksStorage{
		tasks: []*types.Task{
			{
				Id:    1,
				Title: "Task 1",
			},
			{
				Id:    2,
				Title: "Task 2",
			},
		},
	}
}

func (ts *InMemoryTasksStorage) List() ([]*types.Task, error) {
	return ts.tasks, nil
}

func (ts *InMemoryTasksStorage) Update(task *types.Task) error {
	for _, t := range ts.tasks {
		if t.Id == task.Id {
			t.Title = task.Title
		}
	}

	return nil
}

func (ts *InMemoryTasksStorage) Save(task *types.Task) error {
	task.Id = rand.Intn(100)
	ts.tasks = append(ts.tasks, task)

	return nil
}

func (ts *InMemoryTasksStorage) Get(id string) (*types.Task, error) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	for _, task := range ts.tasks {
		if task.Id == intId {
			return task, nil
		}
	}

	return nil, errors.New("Task not found")
}

func (ts *InMemoryTasksStorage) Delete(id string) error {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	for i, task := range ts.tasks {
		if task.Id == intId {
			ts.tasks = append(ts.tasks[:i], ts.tasks[i+1:]...)
			break
		}
	}

	return nil
}
