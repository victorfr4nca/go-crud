package memory

import (
	"errors"
	"math/rand"
	"strconv"
	"sync"

	"github.com/victorfr4nca/go-crud/internal/entity"
	"github.com/victorfr4nca/go-crud/internal/repository/task"
)

type Repository struct {
	tasks []*entity.Task
	sync.Mutex
}

var _ task.Repository = (*Repository)(nil)

func New() *Repository {
	return &Repository{
		tasks: []*entity.Task{
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

func (ts *Repository) List() ([]*entity.Task, error) {
	return ts.tasks, nil
}

func (ts *Repository) Update(task *entity.Task) error {
	ts.Lock()
	for _, t := range ts.tasks {
		if t.Id == task.Id {
			t.Title = task.Title
			return nil
		}
	}
	ts.Unlock()

	return errors.New("Task not found")
}

func (ts *Repository) Save(task *entity.Task) error {
	task.Id = rand.Intn(100)
	ts.Lock()
	ts.tasks = append(ts.tasks, task)
	ts.Unlock()

	return nil
}

func (ts *Repository) Get(id string) (*entity.Task, error) {
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

func (ts *Repository) Delete(id string) error {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	ts.Lock()
	for i, task := range ts.tasks {
		if task.Id == intId {
			ts.tasks = append(ts.tasks[:i], ts.tasks[i+1:]...)
			return nil
		}
	}
	ts.Unlock()

	return errors.New("Task not found")
}
