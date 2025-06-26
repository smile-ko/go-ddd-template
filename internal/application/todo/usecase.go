package application

import (
	"context"
	"time"

	"github.com/google/uuid"
	domain "github.com/smile-ko/go-ddd-template/internal/domain/todo"
)

type todoUsecase struct {
	repo domain.ITodoRepository
}

func NewTodoUseCase(repo domain.ITodoRepository) ITodoUsecase {
	return &todoUsecase{repo: repo}
}

func (uc *todoUsecase) Create(ctx context.Context, in CreateTodoInput) (*TodoOutput, error) {
	t := domain.Todo{
		ID:          uuid.NewString(),
		Title:       in.Title,
		Description: in.Description,
		Status:      "pending",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	created, err := uc.repo.CreateTodo(ctx, t)
	if err != nil {
		return nil, err
	}

	return toOutput(created), nil
}

func (uc *todoUsecase) Get(ctx context.Context, id string) (*TodoOutput, error) {
	todo, err := uc.repo.GetTodoByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return toOutput(todo), nil
}

func (uc *todoUsecase) List(ctx context.Context) ([]*TodoOutput, error) {
	todos, err := uc.repo.ListTodos(ctx)
	if err != nil {
		return nil, err
	}

	var out []*TodoOutput
	for _, t := range todos {
		out = append(out, toOutput(t))
	}
	return out, nil
}

func toOutput(t domain.Todo) *TodoOutput {
	return &TodoOutput{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Status:      t.Status,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}
