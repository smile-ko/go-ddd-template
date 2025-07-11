// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: todo_query.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createTodo = `-- name: CreateTodo :one
INSERT INTO todos (title, description, completed)
VALUES ($1, $2, $3)
RETURNING id, title, description, completed, created_at, updated_at
`

type CreateTodoParams struct {
	Title       string
	Description sql.NullString
	Completed   bool
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) (Todo, error) {
	row := q.db.QueryRow(ctx, createTodo, arg.Title, arg.Description, arg.Completed)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Completed,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteTodo = `-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = $1
`

func (q *Queries) DeleteTodo(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteTodo, id)
	return err
}

const getTodoByID = `-- name: GetTodoByID :one
SELECT id, title, description, completed, created_at, updated_at FROM todos
WHERE id = $1
`

func (q *Queries) GetTodoByID(ctx context.Context, id int32) (Todo, error) {
	row := q.db.QueryRow(ctx, getTodoByID, id)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Completed,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listTodos = `-- name: ListTodos :many
SELECT id, title, description, completed, created_at, updated_at FROM todos
ORDER BY created_at DESC
`

func (q *Queries) ListTodos(ctx context.Context) ([]Todo, error) {
	rows, err := q.db.Query(ctx, listTodos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Completed,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTodo = `-- name: UpdateTodo :one
UPDATE todos
SET title = $2,
    description = $3,
    completed = $4,
    updated_at = NOW()
WHERE id = $1
RETURNING id, title, description, completed, created_at, updated_at
`

type UpdateTodoParams struct {
	ID          int32
	Title       string
	Description sql.NullString
	Completed   bool
}

func (q *Queries) UpdateTodo(ctx context.Context, arg UpdateTodoParams) (Todo, error) {
	row := q.db.QueryRow(ctx, updateTodo,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.Completed,
	)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Completed,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
