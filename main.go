package main

import (
	"Gotry_http/handler"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", handler.PingHandler)
	http.HandleFunc("/hello", handler.HelloHandler)
	http.HandleFunc("/user", handler.UserHandler)

	http.HandleFunc("/users", handler.UsersHandler)

	fmt.Println("server is running at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("start server failed:", err)
	}
}
