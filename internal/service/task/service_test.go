package task

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/victorfr4nca/go-crud/internal/entity"
	"github.com/victorfr4nca/go-crud/mocks"
)

func TestTaskService_List(t *testing.T) {
	t.Run("should return list of tasks", func(t *testing.T) {
		storageMock := &mocks.Repository{}

		ts := New(storageMock)

		storageMock.On("List").Return([]*entity.Task{
			{
				Id:    1,
				Title: "Task 1",
			},
			{
				Id:    2,
				Title: "Task 2",
			},
		}, nil)

		tasks, err := ts.List()
		assert.Nil(t, err)
		assert.NotEmpty(t, tasks)
	})

	t.Run("should return error", func(t *testing.T) {
		storageMock := &mocks.Repository{}

		ts := New(storageMock)

		storageMock.On("List").Return(nil, assert.AnError)

		tasks, err := ts.List()
		assert.Error(t, err)
		assert.Nil(t, tasks)
	})
}

func TestTaskService_Create(t *testing.T) {
	t.Run("should create task", func(t *testing.T) {
		storageMock := &mocks.Repository{}

		ts := New(storageMock)

		task := &entity.Task{
			Title: "Task 1",
		}

		storageMock.On("Save", task).Return(nil)

		createdTask, err := ts.Create(task)
		assert.Nil(t, err)
		assert.NotNil(t, createdTask)
		assert.Equal(t, task.Title, createdTask.Title)
	})

	t.Run("should return error", func(t *testing.T) {
		storageMock := &mocks.Repository{}

		ts := New(storageMock)

		task := &entity.Task{
			Title: "Task 1",
		}

		storageMock.On("Save", task).Return(assert.AnError)

		createdTask, err := ts.Create(task)
		assert.Error(t, err)
		assert.Nil(t, createdTask)
	})
}

func TestTaskService_Update(t *testing.T) {
	t.Run("should update task", func(t *testing.T) {
		storageMock := &mocks.Repository{}

		ts := New(storageMock)

		task := &entity.Task{
			Id:    1,
			Title: "Task 1 Updated",
		}

		storageMock.On("Update", task).Return(nil)

		updatedTask, err := ts.Update(task)
		assert.Nil(t, err)
		assert.NotNil(t, updatedTask)
		assert.Equal(t, task.Title, updatedTask.Title)
	})

	t.Run("should return error", func(t *testing.T) {
		storageMock := &mocks.Repository{}

		ts := New(storageMock)

		task := &entity.Task{
			Id:    1,
			Title: "Task 1",
		}

		storageMock.On("Update", task).Return(assert.AnError)

		updatedTask, err := ts.Update(task)
		assert.Error(t, err)
		assert.Nil(t, updatedTask)
	})
}

func TestTaskService_Delete(t *testing.T) {
	t.Run("should delete task", func(t *testing.T) {
		storageMock := &mocks.Repository{}

		ts := New(storageMock)

		storageMock.On("Delete", "1").Return(nil)

		err := ts.Delete("1")
		assert.Nil(t, err)
	})

	t.Run("should return error", func(t *testing.T) {
		storageMock := &mocks.Repository{}

		ts := New(storageMock)

		storageMock.On("Delete", "1").Return(assert.AnError)

		err := ts.Delete("1")
		assert.Error(t, err)
	})
}
