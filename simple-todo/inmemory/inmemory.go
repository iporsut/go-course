package inmemory

import (
	"simple-todo/todo"
	"time"

	"github.com/google/uuid"
)

type TodoStorage struct {
	todoList []todo.Todo
}

func (t *TodoStorage) GetAll() []todo.Todo {
	return t.todoList
}

func (t *TodoStorage) AddTodo(todo todo.Todo) todo.Todo {
	todo.ID = uuid.NewString()
	todo.CreatedAt = time.Now()
	t.todoList = append(t.todoList, todo)
	return todo
}

func (t *TodoStorage) GetByID(id string) *todo.Todo {
	for _, td := range t.todoList {
		if td.ID == id {
			return &td
		}
	}

	return nil
}
