package application

import "time"

type CreateTodoInput struct {
	Title       string
	Description string
	Completed   bool
}

type TodoOutput struct {
	ID          string
	Title       string
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
