package main

import (
	"Gotry_http/handler"
	"Gotry_http/repo"
	"Gotry_http/service"
	"fmt"
	"net/http"
)

func main() {

	userRepo := repo.NewUserRepo()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	http.HandleFunc("/ping", handler.PingHandler)
	http.HandleFunc("/hello", handler.HelloHandler)
	http.HandleFunc("/user", userHandler.GetUser)
	http.HandleFunc("/users", userHandler.ListUsers)
	http.HandleFunc("/user/create", userHandler.CreateUser)

	fmt.Println("server is running at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("start server failed:", err)
	}
}
