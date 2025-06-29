// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package database

import (
	"context"
)

type Querier interface {
	CreateTodo(ctx context.Context, arg CreateTodoParams) (Todo, error)
	DeleteTodo(ctx context.Context, id int32) error
	GetTodoByID(ctx context.Context, id int32) (Todo, error)
	ListTodos(ctx context.Context) ([]Todo, error)
	UpdateTodo(ctx context.Context, arg UpdateTodoParams) (Todo, error)
}

var _ Querier = (*Queries)(nil)
