-- Todoタスクを管理するためのSQLクエリ

-- Todoタスクを作成
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