package storage

import "github.com/victorfr4nca/go-crud/internal/types"

type TasksStorage interface {
	List() ([]*types.Task, error)
	Get(id string) (*types.Task, error)
	Update(task *types.Task) error
	Save(task *types.Task) error
	Delete(id string) error
}
