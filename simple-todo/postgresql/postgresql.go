package postgresql

import (
	"context"
	"database/sql"
	"simple-todo/todo"

	"github.com/google/uuid"
)

type DB struct {
	*sql.DB
}

func NewDB(db *sql.DB) *DB {
	return &DB{db}
}

func (db *DB) GetAll(ctx context.Context) ([]*todo.Todo, error) {
	rows, err := db.QueryContext(ctx, "SELECT * FROM todos")
	if err != nil {
		return nil, err
	}

	todos := []*todo.Todo{}
	for rows.Next() {
		var todo todo.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.CreatedAt, &todo.CompletedAt); err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func (db *DB) AddTodo(ctx context.Context, todo *todo.Todo) (*todo.Todo, error) {
	todo.ID = uuid.NewString()

	_, err := db.ExecContext(ctx, "INSERT INTO todos (id, title, created_at) VALUES ($1, $2, CURRENT_TIMESTAMP)", todo.ID, todo.Title)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (db *DB) GetByID(ctx context.Context, id string) (*todo.Todo, error) {
	row := db.QueryRowContext(ctx, "SELECT * FROM todos WHERE id = $1", id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var todo todo.Todo
	if err := row.Scan(&todo.ID, &todo.Title, &todo.CreatedAt, &todo.CompletedAt); err != nil {
		return nil, err
	}

	return &todo, nil
}
