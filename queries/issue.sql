-- name: CreateIssueTask :exec

INSERT INTO issues (id, status, solution, cause)
VALUES (sqlc.arg(id),
        sqlc.arg(status),
        sqlc.arg(solution),
        sqlc.arg(cause));

-- Issueタスクを取得（tasksとissuesをJOINして取得）
-- name: GetIssueTask :one

SELECT t.id,
       t.title,
       t.description,
       t.created_at,
       t.updated_at,
       i.status,
       i.solution,
       i.cause
FROM tasks t
INNER JOIN issues i ON t.id = i.id
WHERE t.id = sqlc.arg(id); -- 全てのIssueタスクを取得
-- name: ListIssueTasks :many

SELECT t.id,
       t.title,
       t.description,
       t.created_at,
       t.updated_at,
       i.status,
       i.solution,
       i.cause
FROM tasks t
INNER JOIN issues i ON t.id = i.id
ORDER BY t.created_at DESC;

-- Issueを更新
-- name: UpdateIssueStatus :exec

UPDATE issues
SET status = sqlc.arg(status),
    solution = sqlc.arg(solution),
    cause = sqlc.arg(cause)
WHERE id = sqlc.arg(id);

-- Issueタスクを削除（外部キー制約により、issuesも自動削除される）
-- name: DeleteIssueTask :exec

DELETE
FROM tasks
WHERE id = sqlc.arg(id);