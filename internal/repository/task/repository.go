package task

import "github.com/victorfr4nca/go-crud/internal/entity"

type Repository interface {
	List() ([]*entity.Task, error)
	Get(id string) (*entity.Task, error)
	Update(task *entity.Task) error
	Save(task *entity.Task) error
	Delete(id string) error
}
