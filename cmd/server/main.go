package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/leona-art/task-manager/adaptor/controller"
	"github.com/leona-art/task-manager/gen/workspace/v1/workspacev1connect"
	"github.com/leona-art/task-manager/infra/mysql_repository"
	"github.com/leona-art/task-manager/usecase"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "app_user:app_password@/app_db?parseTime=true")
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}
	if err := db.Ping(); err != nil {
		panic(fmt.Sprintf("failed to ping database: %v", err))
	}
	defer db.Close()

	todoRepository := mysql_repository.NewMySqlTodoRepository(db)
	todoUsecase := usecase.NewTodoUseCase(todoRepository)
	workspace := &controller.WorkSpaceController{
		TodoUsecase: todoUsecase,
	}
	mux := http.NewServeMux()
	path, handler := workspacev1connect.NewWorkspaceServiceHandler(workspace)
	mux.Handle(path, handler)
	http.ListenAndServe(
		":8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
