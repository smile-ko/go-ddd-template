package application

import "context"

type ITodoUsecase interface {
	Create(ctx context.Context, in CreateTodoInput) (*TodoOutput, error)
	Get(ctx context.Context, id string) (*TodoOutput, error)
	List(ctx context.Context) ([]*TodoOutput, error)
}
