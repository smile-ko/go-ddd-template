package domain

import (
	"context"
)

type ITodoRepository interface {
	CreateTodo(ctx context.Context, arg Todo) (Todo, error)
	GetTodoByID(ctx context.Context, id string) (Todo, error)
	ListTodos(ctx context.Context) ([]Todo, error)
}
