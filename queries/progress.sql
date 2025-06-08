
-- name: CreateProgressTask :exec

INSERT INTO progress (id, status, solution)
VALUES (sqlc.arg(id),
        sqlc.arg(status),
        sqlc.arg(solution));

-- Progressタスクを取得（tasksとprogressをJOINして取得）
-- name: GetProgressTask :one

SELECT t.id,
       t.title,
       t.description,
       t.created_at,
       t.updated_at,
       p.status,
       p.solution
FROM tasks t
INNER JOIN progress p ON t.id = p.id
WHERE t.id = sqlc.arg(id);

-- 全てのProgressタスクを取得
-- name: ListProgressTasks :many

SELECT t.id,
       t.title,
       t.description,
       t.created_at,
       t.updated_at,
       p.status,
       p.solution
FROM tasks t
INNER JOIN progress p ON t.id = p.id
ORDER BY t.created_at DESC; -- Progressを更新
-- name: UpdateProgressStatus :exec

UPDATE progress
SET status = sqlc.arg(status),
    solution = sqlc.arg(solution)
WHERE id = sqlc.arg(id);

-- Progressタスクを削除（外部キー制約により、progressも自動削除される）
-- name: DeleteProgressTask :exec

DELETE
FROM tasks
WHERE id = sqlc.arg(id);
