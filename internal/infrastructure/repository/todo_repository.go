package repository

import (
	"context"
	"database/sql"
	"strconv"

	domain "github.com/smile-ko/go-ddd-template/internal/domain/todo"
	"github.com/smile-ko/go-ddd-template/internal/infrastructure/db/sqlc"
)

type sqlTodoRepository struct {
	q *sqlc.Queries
}

func NewTodoRepository(q *sqlc.Queries) domain.ITodoRepository {
	return &sqlTodoRepository{q: q}
}

func (r *sqlTodoRepository) CreateTodo(ctx context.Context, t domain.Todo) (domain.Todo, error) {
	arg := sqlc.CreateTodoParams{
		Title: t.Title,
		Description: sql.NullString{
			String: t.Description,
			Valid:  t.Description != "",
		},
		Completed: t.Status == "completed",
	}
	created, err := r.q.CreateTodo(ctx, arg)
	if err != nil {
		return domain.Todo{}, err
	}
	return toDomain(created), nil
}

func (r *sqlTodoRepository) GetTodoByID(ctx context.Context, id string) (domain.Todo, error) {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return domain.Todo{}, err
	}

	dbTodo, err := r.q.GetTodoByID(ctx, int32(intID))
	if err != nil {
		return domain.Todo{}, err
	}

	return toDomain(dbTodo), nil
}

func (r *sqlTodoRepository) ListTodos(ctx context.Context) ([]domain.Todo, error) {
	dbTodos, err := r.q.ListTodos(ctx)
	if err != nil {
		return nil, err
	}

	var todos []domain.Todo
	for _, dbTodo := range dbTodos {
		todos = append(todos, toDomain(dbTodo))
	}
	return todos, nil
}

func toDomain(t sqlc.Todo) domain.Todo {
	return domain.Todo{
		ID:          strconv.Itoa(int(t.ID)),
		Title:       t.Title,
		Description: t.Description.String,
		Status:      boolToStatus(t.Completed),
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}

func boolToStatus(done bool) string {
	if done {
		return "completed"
	}
	return "pending"
}
