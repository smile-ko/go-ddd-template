// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package sqlc

import (
	"database/sql"
	"time"
)

type Todo struct {
	ID          int32
	Title       string
	Description sql.NullString
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
