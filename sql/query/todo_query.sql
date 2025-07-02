
-- name: CreateTodo :one
INSERT INTO todos (title, description, completed)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetTodoByID :one
SELECT * FROM todos
WHERE id = $1;

-- name: ListTodos :many
SELECT * FROM todos
ORDER BY created_at DESC;

-- name: UpdateTodo :one
UPDATE todos
SET title = $2,
    description = $3,
    completed = $4,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = $1;
