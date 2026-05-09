package main

import (
	"Gotry_http/db"
	"Gotry_http/handler"
	"Gotry_http/repo"
	"Gotry_http/service"
	"fmt"
	"net/http"
)

func main() {

	mysqlDB, err := db.NewMySQLDB()
	if err != nil {
		panic(err)
	}

	// user 模块
	userRepo := repo.NewUserRepo(mysqlDB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// task 模块
	taskRepo := repo.NewTaskRepo(mysqlDB)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	http.HandleFunc("/ping", handler.PingHandler)
	http.HandleFunc("/hello", handler.HelloHandler)

	http.HandleFunc("/user", userHandler.GetUser)
	http.HandleFunc("/users", userHandler.ListUsers)
	http.HandleFunc("/user/create", userHandler.CreateUser)

	http.HandleFunc("/tasks", taskHandler.Tasks)
	http.HandleFunc("/task", taskHandler.TaskDetail)
	http.HandleFunc("/task/status", taskHandler.UpdateTaskStatus)

	http.HandleFunc("/debug/async-test", taskHandler.AsyncTestHandler)

	fmt.Println("server is running at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("start server failed:", err)
	}
}
