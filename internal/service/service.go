package service

import "github.com/victorfr4nca/go-crud/internal/types"

type TaskService interface {
	List() ([]*types.Task, error)
	Create(task *types.Task) (*types.Task, error)
	Update(task *types.Task) (*types.Task, error)
	Delete(id string) error
}
