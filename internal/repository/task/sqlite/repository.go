package sqlite

import (
	"database/sql"
	"sync"

	"github.com/victorfr4nca/go-crud/internal/entity"
	"github.com/victorfr4nca/go-crud/internal/repository/task"
)

type Repository struct {
	db *sql.DB
	sync.Mutex
}

var _ task.Repository = (*Repository)(nil)

func New(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (ts *Repository) List() ([]*entity.Task, error) {
	rows, err := ts.db.Query(
		"SELECT * FROM tasks",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data := []*entity.Task{}
	for rows.Next() {
		i := &entity.Task{}
		err = rows.Scan(&i.Id, &i.Title)
		if err != nil {
			return nil, err
		}
		data = append(data, i)
	}

	return data, nil
}

func (ts *Repository) Update(task *entity.Task) error {
	return nil
}

func (ts *Repository) Save(task *entity.Task) error {
	tx, err := ts.db.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.Exec("INSERT INTO tasks VALUES(NULL,?);", task.Title)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (ts *Repository) Get(id string) (*entity.Task, error) {
	row := ts.db.QueryRow(
		"SELECT * FROM tasks WHERE id = ?",
		id,
	)

	data := &entity.Task{}
	err := row.Scan(&data.Id, &data.Title)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (ts *Repository) Delete(id string) error {
	return nil
}
