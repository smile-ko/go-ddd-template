package domain

import "time"

type Status string

const (
	Pending    Status = "pending"
	InProgress Status = "in_progress"
	Done       Status = "done"
)

type Todo struct {
	ID          string
	Title       string
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
