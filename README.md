## ユビキタス

- タスク
    - Todo: Todoタスク
    - Progress: 進捗管理タスク
    - Issue: 課題管理タスク

```bash
atlas schema apply -u "mysql://root:root@127.0.0.1:3306/app_db" --to file://schema.sql --dev-url "docker://mysql/8/app_db"
```