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

-- name: CreateTodoTask :exec

INSERT INTO todos (id, status)
VALUES (sqlc.arg(id),
        sqlc.arg(status));

-- Todoタスクを取得（tasksとtodosをJOINして取得）
-- name: GetTodoTask :one

SELECT t.id,
       t.title,
       t.description,
       t.created_at,
       t.updated_at,
       td.status
FROM tasks t
INNER JOIN todos td ON t.id = td.id
WHERE t.id = sqlc.arg(id);

-- 全てのTodoタスクを取得
-- name: ListTodoTasks :many

SELECT t.id,
       t.title,
       t.description,
       t.created_at,
       t.updated_at,
       td.status
FROM tasks t
INNER JOIN todos td ON t.id = td.id
ORDER BY t.created_at DESC;

-- Todoのステータスを更新
-- name: UpdateTodoStatus :exec

UPDATE todos
SET status = sqlc.arg(status)
WHERE id = sqlc.arg(id);

-- Todoタスクを削除（外部キー制約により、todosも自動削除される）
-- name: DeleteTodoTask :exec

DELETE
FROM tasks
WHERE id = sqlc.arg(id);

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