package task

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/victorfr4nca/go-crud/internal/entity"
	"github.com/victorfr4nca/go-crud/mocks"
)

func TestTaskHandler_GetHandler(t *testing.T) {
	t.Run("should return list of tasks", func(t *testing.T) {
		serviceMock := &mocks.Service{}

		th := New(serviceMock)
		req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
		w := httptest.NewRecorder()

		serviceMock.On("List").Return([]*entity.Task{
			{
				Id:    1,
				Title: "Task 1",
			},
		}, nil)

		th.GetHandler(w, req)

		res := w.Result()
		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)

		assert.Nil(t, err)
		assert.NotNil(t, data)
		assert.Contains(t, string(data), "Task 1")
	})

	t.Run("should return error", func(t *testing.T) {
		serviceMock := &mocks.Service{}

		th := New(serviceMock)
		req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
		w := httptest.NewRecorder()

		serviceMock.On("List").Return(nil, assert.AnError)

		th.GetHandler(w, req)

		res := w.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	})
}
