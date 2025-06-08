-- name: CreateTask :exec

INSERT INTO tasks (id, title, description,kind, created_at, updated_at)
VALUES (sqlc.arg(id),
        sqlc.arg(title),
        sqlc.arg(description),
        sqlc.arg(kind),
        sqlc.arg(created_at),
        sqlc.arg(updated_at));

-- name: UpdateTask :exec

UPDATE tasks
SET title = sqlc.arg(title),
    description = sqlc.arg(description),
    kind = sqlc.arg(kind),
    updated_at = sqlc.arg(updated_at)
WHERE id = sqlc.arg(id);

-- name: DeleteTask :exec

DELETE
FROM tasks
WHERE id = sqlc.arg(id);