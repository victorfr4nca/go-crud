package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/victorfr4nca/go-crud/internal/entity"
)

func TestInMemoryTasksStorage_New(t *testing.T) {
	t.Run("should return a new task service", func(t *testing.T) {
		taskService := New()
		assert.NotNil(t, taskService)
	})

	t.Run("tasks should not be empty", func(t *testing.T) {
		taskService := New()
		assert.NotNil(t, taskService)
		assert.NotNil(t, taskService.tasks)
		assert.Equal(t, 2, len(taskService.tasks))
		assert.Equal(t, "Task 1", taskService.tasks[0].Title)
		assert.Equal(t, "Task 2", taskService.tasks[1].Title)
	})
}

func TestInMemoryTasksStorage_List(t *testing.T) {
	t.Run("should return a list of tasks", func(t *testing.T) {
		taskService := New()
		tasks, err := taskService.List()
		assert.Nil(t, err)
		assert.NotNil(t, tasks)
		assert.Equal(t, 2, len(tasks))
		assert.Equal(t, "Task 1", taskService.tasks[0].Title)
		assert.Equal(t, "Task 2", taskService.tasks[1].Title)
	})
}

func TestInMemoryTasksStorage_Update(t *testing.T) {
	t.Run("should update a task", func(t *testing.T) {
		taskService := New()
		task := &entity.Task{Id: 1, Title: "Task 1 Updated"}
		err := taskService.Update(task)
		assert.Nil(t, err)
		assert.Equal(t, "Task 1 Updated", taskService.tasks[0].Title)
	})

	t.Run("should return an error when task does not exist", func(t *testing.T) {
		taskService := New()
		task := &entity.Task{Id: 3, Title: "Task 1 Updated"}
		err := taskService.Update(task)
		assert.NotNil(t, err)
		assert.Error(t, err, "Task not found")
	})
}

func TestInMemoryTasksStorage_Save(t *testing.T) {
	t.Run("should save a task", func(t *testing.T) {
		taskService := New()
		task := &entity.Task{Title: "Task 3"}
		err := taskService.Save(task)
		assert.Nil(t, err)
		assert.Equal(t, 3, len(taskService.tasks))
		assert.Equal(t, "Task 3", taskService.tasks[2].Title)
	})
}

func TestInMemoryTasksStorage_Get(t *testing.T) {
	t.Run("should get a task", func(t *testing.T) {
		taskService := New()
		task, err := taskService.Get("1")
		assert.Nil(t, err)
		assert.NotNil(t, task)
		assert.Equal(t, "Task 1", task.Title)
	})

	t.Run("should return an error when task does not exist", func(t *testing.T) {
		taskService := New()
		task, err := taskService.Get("3")
		assert.NotNil(t, err)
		assert.Nil(t, task)
		assert.Error(t, err, "Task not found")
	})
}

func TestInMemoryTasksStorage_Delete(t *testing.T) {
	t.Run("should delete a task", func(t *testing.T) {
		taskService := New()
		err := taskService.Delete("1")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(taskService.tasks))
		assert.NotContains(t, taskService.tasks, &entity.Task{Id: 1, Title: "Task 1"})
	})

	t.Run("should return an error when task does not exist", func(t *testing.T) {
		taskService := New()
		err := taskService.Delete("3")
		assert.NotNil(t, err)
		assert.Error(t, err, "Task not found")
	})
}
